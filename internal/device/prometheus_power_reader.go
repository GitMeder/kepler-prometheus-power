// SPDX-FileCopyrightText: 2025 The Kepler Authors
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// PrometheusClient executes instant queries against Prometheus-compatible APIs.
type PrometheusClient struct {
	baseURL  string
	query    string
	username string
	password string
	client   *http.Client
	logger   *slog.Logger
}

// NewPrometheusClient builds a simple client for issuing power queries.
func NewPrometheusClient(baseURL, query, caFile, username, passwordFile string, logger *slog.Logger) (*PrometheusClient, error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	if strings.TrimSpace(caFile) != "" {
		caPEM, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read prometheus CA file %q: %w", caFile, err)
		}

		roots, err := x509.SystemCertPool()
		if err != nil || roots == nil {
			roots = x509.NewCertPool()
		}

		if ok := roots.AppendCertsFromPEM(caPEM); !ok {
			return nil, fmt.Errorf("failed to append CA certificates from %q", caFile)
		}

		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{
			RootCAs: roots,
		}

		httpClient.Transport = transport
	}

	password := ""
	if strings.TrimSpace(passwordFile) != "" {
		raw, err := os.ReadFile(passwordFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read prometheus password file %q: %w", passwordFile, err)
		}
		password = strings.TrimSpace(string(raw))
	}

	if logger != nil {
		logger = logger.With("component", "prometheus-power-reader")
	}

	return &PrometheusClient{
		baseURL:  strings.TrimRight(baseURL, "/"),
		query:    query,
		username: strings.TrimSpace(username),
		password: password,
		client:   httpClient,
		logger:   logger,
	}, nil
}

// QueryPowerWatt executes the configured query and returns the watt value.
func (c *PrometheusClient) QueryPowerWatt(ctx context.Context) (float64, error) {
	started := time.Now()

	u, err := url.Parse(c.baseURL + "/api/v1/query")
	if err != nil {
		return 0, fmt.Errorf("invalid base URL %q: %w", c.baseURL, err)
	}

	q := u.Query()
	q.Set("query", c.query)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to build prometheus query request: %w", err)
	}

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	if c.logger != nil {
		c.logger.Info(
			"Starting Prometheus power query",
			"query", c.query,
			"url", u.String(),
		)
	}

	resp, err := c.client.Do(req)
	requestDuration := time.Since(started)
	if err != nil {
		if c.logger != nil {
			c.logger.Error(
				"Prometheus power query failed",
				"query", c.query,
				"duration", requestDuration,
				"error", err,
			)
		}
		return 0, fmt.Errorf("failed to execute prometheus query after %s: %w", requestDuration, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if c.logger != nil {
			c.logger.Error(
				"Prometheus power query returned non-200 status",
				"query", c.query,
				"duration", requestDuration,
				"status", resp.Status,
			)
		}
		return 0, fmt.Errorf("prometheus query failed after %s: %s", requestDuration, resp.Status)
	}

	var result promQueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		if c.logger != nil {
			c.logger.Error(
				"Failed to decode Prometheus power query response",
				"query", c.query,
				"duration", time.Since(started),
				"error", err,
			)
		}
		return 0, fmt.Errorf("failed to decode prometheus response: %w", err)
	}

	if result.Status != "success" {
		if c.logger != nil {
			c.logger.Error(
				"Prometheus power query returned non-success result",
				"query", c.query,
				"duration", time.Since(started),
				"status", result.Status,
			)
		}
		return 0, fmt.Errorf("prometheus response status: %s", result.Status)
	}

	if len(result.Data.Result) == 0 {
		if c.logger != nil {
			c.logger.Warn(
				"Prometheus power query returned no samples",
				"query", c.query,
				"duration", time.Since(started),
			)
		}
		return 0, fmt.Errorf("prometheus query %q returned no samples", c.query)
	}

	v := result.Data.Result[0].Value[1]

	switch t := v.(type) {
	case string:
		watts, err := strconv.ParseFloat(t, 64)
		if err != nil {
			if c.logger != nil {
				c.logger.Error(
					"Failed to parse Prometheus power value",
					"query", c.query,
					"duration", time.Since(started),
					"value", t,
					"error", err,
				)
			}
			return 0, fmt.Errorf("failed to parse prometheus value %q as float64: %w", t, err)
		}
		if c.logger != nil {
			c.logger.Info(
				"Prometheus power query succeeded",
				"query", c.query,
				"duration", time.Since(started),
				"watts", watts,
			)
		}
		return watts, nil
	case float64:
		if c.logger != nil {
			c.logger.Info(
				"Prometheus power query succeeded",
				"query", c.query,
				"duration", time.Since(started),
				"watts", t,
			)
		}
		return t, nil
	default:
		if c.logger != nil {
			c.logger.Error(
				"Unexpected Prometheus response value type",
				"query", c.query,
				"duration", time.Since(started),
				"type", fmt.Sprintf("%T", v),
			)
		}
		return 0, fmt.Errorf("unexpected value type %T in prometheus response", v)
	}
}

type promQueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  [2]any            `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
