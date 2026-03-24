// SPDX-FileCopyrightText: 2025 The Kepler Authors
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

// PrometheusPowerMeter implements CPUPowerMeter backed by an external power feed.
type PrometheusPowerMeter struct {
	zone *PrometheusPowerZone
}

// PrometheusPowerZone exposes a single power-only zone (no energy counters).
type PrometheusPowerZone struct {
	name      string
	index     int
	path      string
	readPower func(ctx context.Context) (float64, error)
	logger    *slog.Logger
}

// Name returns the meter identifier.
func (m *PrometheusPowerMeter) Name() string {
	return "prometheus-power-meter"
}

// Zones returns the single zone exposed by this meter.
func (m *PrometheusPowerMeter) Zones() ([]EnergyZone, error) {
	return []EnergyZone{m.zone}, nil
}

// PrimaryEnergyZone returns the only zone exposed by this meter.
func (m *PrometheusPowerMeter) PrimaryEnergyZone() (EnergyZone, error) {
	return m.zone, nil
}

// Start is a no-op because the meter only proxies Prometheus data.
func (m *PrometheusPowerMeter) Start() error { return nil }

// Stop is a no-op because the meter only proxies Prometheus data.
func (m *PrometheusPowerMeter) Stop() error { return nil }

// Name returns the zone identifier.
func (z *PrometheusPowerZone) Name() string { return z.name }

// Index always returns zero because the meter only exposes a single zone.
func (z *PrometheusPowerZone) Index() int { return z.index }

// Path returns the logical path used for the Prometheus power feed.
func (z *PrometheusPowerZone) Path() string { return z.path }

// Energy returns zero because the Prometheus feed only reports instantaneous power.
func (z *PrometheusPowerZone) Energy() (Energy, error) { return 0, nil }

// MaxEnergy returns zero to signal that this is a power-only sensor.
func (z *PrometheusPowerZone) MaxEnergy() Energy { return 0 }

// Power queries the external feed and converts the watt value to microwatts.
func (z *PrometheusPowerZone) Power() (Power, error) {
	started := time.Now()

	if z.logger != nil {
		z.logger.Info(
			"Starting Prometheus power read",
			"zone", z.name,
			"path", z.path,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	watts, err := z.readPower(ctx)
	duration := time.Since(started)
	if err != nil {
		if z.logger != nil {
			z.logger.Error(
				"Prometheus power read failed",
				"zone", z.name,
				"path", z.path,
				"duration", duration,
				"error", err,
			)
		}
		return 0, fmt.Errorf("prometheus power read failed after %s: %w", duration, err)
	}

	if watts < 0 {
		if z.logger != nil {
			z.logger.Error(
				"Prometheus power read returned negative watt value",
				"zone", z.name,
				"path", z.path,
				"duration", duration,
				"watts", watts,
			)
		}
		return 0, fmt.Errorf("received negative watt value %.2f from Prometheus after %s", watts, duration)
	}

	if z.logger != nil {
		z.logger.Info(
			"Prometheus power read succeeded",
			"zone", z.name,
			"path", z.path,
			"duration", duration,
			"watts", watts,
		)
	}

	return Power(watts * 1_000_000), nil
}

// NewPrometheusPowerMeter builds a CPUPowerMeter backed by a Prometheus/VictoriaMetrics query.
func NewPrometheusPowerMeter(
	zoneName string,
	readPowerFunc func(ctx context.Context) (float64, error),
	logger *slog.Logger,
) (CPUPowerMeter, error) {
	if readPowerFunc == nil {
		return nil, fmt.Errorf("readPowerFunc must not be nil")
	}

	name := strings.TrimSpace(zoneName)
	if name == "" {
		name = "prometheus-node"
	}

	if logger != nil {
		logger = logger.With("component", "prometheus-power-meter", "zone", name)
	}

	return &PrometheusPowerMeter{
		zone: &PrometheusPowerZone{
			name:      name,
			index:     0,
			path:      fmt.Sprintf("prometheus:%s", name),
			readPower: readPowerFunc,
			logger:    logger,
		},
	}, nil
}
