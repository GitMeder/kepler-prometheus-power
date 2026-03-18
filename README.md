# Kepler

[![GitHub license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/sustainable-computing-io/kepler/blob/main/LICENSES) [![codecov](https://codecov.io/gh/sustainable-computing-io/kepler/branch/main/graph/badge.svg?token=K9BDX9M86E)](https://codecov.io/gh/sustainable-computing-io/kepler/tree/main) [![CI Status](https://github.com/sustainable-computing-io/kepler/actions/workflows/push.yaml/badge.svg?branch=main)](https://github.com/sustainable-computing-io/kepler/actions/workflows/push.yaml) [![Releases](https://img.shields.io/github/v/tag/sustainable-computing-io/kepler)](https://github.com/sustainable-computing-io/kepler/releases) [![zread](https://img.shields.io/badge/Ask_Zread-_.svg?style=flat&color=00b0aa&labelColor=000000&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIHZpZXdCb3g9IjAgMCAxNiAxNiIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTQuOTYxNTYgMS42MDAxSDIuMjQxNTZDMS44ODgxIDEuNjAwMSAxLjYwMTU2IDEuODg2NjQgMS42MDE1NiAyLjI0MDFWNC45NjAxQzEuNjAxNTYgNS4zMTM1NiAxLjg4ODEgNS42MDAxIDIuMjQxNTYgNS42MDAxSDQuOTYxNTZDNS4zMTUwMiA1LjYwMDEgNS42MDE1NiA1LjMxMzU2IDUuNjAxNTYgNC45NjAxVjIuMjQwMUM1LjYwMTU2IDEuODg2NjQgNS4zMTUwMiAxLjYwMDEgNC45NjE1NiAxLjYwMDFaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00Ljk2MTU2IDEwLjM5OTlIMi4yNDE1NkMxLjg4ODEgMTAuMzk5OSAxLjYwMTU2IDEwLjY4NjQgMS42MDE1NiAxMS4wMzk5VjEzLjc1OTlDMS42MDE1NiAxNC4xMTM0IDEuODg4MSAxNC4zOTk5IDIuMjQxNTYgMTQuMzk5OUg0Ljk2MTU2QzUuMzE1MDIgMTQuMzk5OSA1LjYwMTU2IDE0LjExMzQgNS42MDE1NiAxMy43NTk5VjExLjAzOTlDNS42MDE1NiAxMC42ODY0IDUuMzE1MDIgMTAuMzk5OSA0Ljk2MTU2IDEwLjM5OTlaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik0xMy43NTg0IDEuNjAwMUgxMS4wMzg0QzEwLjY4NSAxLjYwMDEgMTAuMzk4NCAxLjg4NjY0IDEwLjM5ODQgMi4yNDAxVjQuOTYwMUMxMC4zOTg0IDUuMzEzNTYgMTAuNjg1IDUuNjAwMSAxMS4wMzg0IDUuNjAwMUgxMy43NTg0QzE0LjExMTkgNS42MDAxIDE0LjM5ODQgNS4zMTM1NiAxNC4zOTg0IDQuOTYwMVYyLjI0MDFDMTQuMzk4NCAxLjg4NjY0IDE0LjExMTkgMS42MDAxIDEzLjc1ODQgMS42MDAxWiIgZmlsbD0iI2ZmZiIvPgo8cGF0aCBkPSJNNCAxMkwxMiA0TDQgMTJaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00IDEyTDEyIDQiIHN0cm9rZT0iI2ZmZiIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIvPgo8L3N2Zz4K&logoColor=ffffff)](https://zread.ai/sustainable-computing-io/kepler) [![OpenSSF Best Practices](https://www.bestpractices.dev/projects/7391/badge)](https://www.bestpractices.dev/projects/7391) [![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/sustainable-computing-io/kepler/badge)](https://scorecard.dev/viewer/?uri=github.com/sustainable-computing-io/kepler)

Kepler (Kubernetes-based Efficient Power Level Exporter) is a Prometheus exporter that measures energy consumption metrics at the container, pod, and node level in Kubernetes clusters.

## 🚀 Major Rewrite: Kepler (0.10.0 and above)

**Important Notice:** Starting with version 0.10.0, Kepler has undergone a complete ground-up rewrite.
This represents a significant architectural improvement while maintaining the core mission of
accurate energy consumption monitoring for cloud-native workloads.

> 📢 **Read the full announcement:** [CNCF Slack Announcement](https://cloud-native.slack.com/archives/C05QK3KN3HT/p1752049660866519)

### ✨ What's New in the Rewrite

**Enhanced Performance & Accuracy:**

- Dynamic detection of Nodes' RAPL zones - no more hardcoded RAPL zones
- More accurate power attribution based on active CPU usage (no more idle/dynamic for workloads)
- Improved VM, Container, and Pod detection with more meaningful label values
- Significantly reduced resource usage compared to old Kepler

**Reduced Security Requirements:**

- Requires only readonly access to host `/proc` and `/sys`
- No more `CAP_SYSADMIN` or `CAP_BPF` capabilities required
- Much fewer privileges than previous versions

**Modern Architecture:**

- Service-oriented design with clean separation of concerns
- Thread-safe operations throughout the codebase
- Graceful shutdown handling with proper resource cleanup
- Comprehensive error handling with structured logging

**Current Limitations:**

- Only supports Baremetal (platform power support in roadmap)
- Supports only RAPL/powercap framework
- No GPU power support yet

### 📚 Migration & Legacy Support

**For New Users:** Use the current version (0.10.0+) for the best experience and latest features.

**For Existing Users:** If you need to continue using the old version:

- Pin your deployment to version `0.9.0` (final legacy release)
- Access the old codebase in the [archived branch](https://github.com/sustainable-computing-io/kepler/tree/archived)
- **Important:** The legacy version (0.9.x and earlier) is now frozen - no bug fixes or feature requests will be accepted for the old version

**Migration Note:** Please review the new configuration format and deployment methods below when upgrading to 0.10.0+.

## 🚀 Getting Started

> **📖 For comprehensive installation instructions, troubleshooting, and advanced deployment options, see our [Installation Guide](docs/user/installation.md)**

### ⚡ Quick Start (Kubernetes with Helm)

This quick start shows the production deployment path for this custom fork (not the generic upstream OCI chart).

```sh
# 1. Deploy or upgrade Kepler from local fork chart with production values
helm upgrade --install kepler manifests/helm/kepler \
  -f manifests/helm/kepler/values-prompower-prod.yaml \
  --namespace default

# Wait for Kepler pods to be running
kubectl wait --for=condition=ready --timeout=120s pod -n default -l app.kubernetes.io/name=kepler --all

# 2. Verify installation
kubectl get pods -n default -l app.kubernetes.io/name=kepler

# 3. Access metrics (port-forward)
kubectl port-forward -n default svc/kepler 28282:28282

# Test metrics endpoint
curl http://localhost:28282/metrics | grep kepler_node_cpu_watts
```

> **📋 For Production Deployments:** This fork uses a custom local chart and fork image. See the production deployment section for concrete values and secret names.

**Next Steps:**

To ensure Kepler is working correctly and to visualize the metrics:

- **[Verify Metrics Collection](docs/user/installation.md#verify-metrics-collection)** - Verify power consumption metrics are being collected
- **[Configuration Options](docs/user/configuration.md)** - Customize Kepler deployment
- **[Helm Updates & Management](docs/user/helm-updates.md)** - Learn how to upgrade and manage Kepler with Helm

**Need Help?**

- [Installation Guide](docs/user/installation.md) - Detailed prerequisites, configuration options, and installation steps
- [Metrics Documentation](docs/user/metrics.md) - Available metrics and their descriptions

### 🔧 Other Installation Methods (generic upstream references)

Choose your preferred method. The following examples are generic Kepler references and are not the primary production flow for this fork.

```bash
# 💻 Local Development
make build && sudo ./bin/kepler

# ✨ Docker Compose (with Prometheus & Grafana)
cd compose/dev && docker compose up -d

# 🐳 Kubernetes with Kustomize
kubectl kustomize manifests/k8s | \
  sed -e "s|<KEPLER_IMAGE>|quay.io/sustainable_computing_io/kepler:latest|g" | \
  kubectl apply --server-side --force-conflicts -f -
```

## Production Deployment with External Power Input

This fork of Kepler is configured for production use with an external Prometheus/VictoriaMetrics power source and a custom image/chart path. Upstream generic examples are retained for reference, but the actual runtime setup for this repository is concrete and opinionated.

In this deployment, POWER_DEVICE is resolved from `NODE_NAME` by default, and mapping via `nodeDeviceMapFile` is optional fallback.

Kepler can run in two modes:

- Direct mode (default): `POWER_DEVICE = NODE_NAME`
- Optional mapping mode: `nodeDeviceMapFile` supplied to map node names to external device labels when they differ

### External Power Input Modes

- Direct mode: `POWER_DEVICE` resolves to the Kubernetes node name (`NODE_NAME`) by default.
- Optional mapping mode: if `nodeDeviceMapFile` is configured and non-empty, Kepler loads a node-to-device mapping from that file.

### Architecture Overview

- Kepler runs as a DaemonSet on Kubernetes worker nodes.
- An external Prometheus/VictoriaMetrics endpoint provides node/device power input.
- Node-to-device mapping is optional and only required when node names differ from external `device` label values.
- A `vmagent` sidecar scrapes local Kepler metrics and forwards them by `remote_write` to the central monitoring backend.
- Central endpoint is secured with HTTPS + Basic Auth.

### External Power Query

The external power source uses a query template like:

```yaml
query: device_power_watts_avg{room="R3.033",rack="Rack 3",device="${POWER_DEVICE}"}
```

### Required Kubernetes Objects

In this production fork, ensure the following objects are present:

- `gitlab-registry-creds`: image pull secret for the private GitLab registry.
- `vm-ca-cert`: CA secret for TLS trust of the central VictoriaMetrics endpoint.
- `prompower-auth`: secret for Kepler external power query credentials.
- `vm-remote-write-auth`: secret for `vmagent` remote_write credentials.
- Optional `kepler-node-device-map`: ConfigMap for node-to-device mapping, only required if node names differ from external `device` labels.

### Example Node-to-Device ConfigMap (Optional Fallback Mode)

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kepler-node-device-map
  namespace: default
data:
  worker-node-01: kub01
  worker-node-02: kub02
  worker-node-03: kub03
```

### Helm Values for Production

```yaml
namespace:
  create: false
  name: default

serviceAccount:
  create: true
  name: kepler

image:
  repository: gitlab.lrz.de:5005/green-it/students/bachelor/eder-moritz/kepler-prompower
  tag: v0.1.2
  pullPolicy: IfNotPresent
imagePullSecrets:
  - name: gitlab-registry-creds

config:
  experimental:
    prometheus-power:
      enabled: true
      baseURL: "https://prometheus.cs.hm.edu:8428"
      username: "vmuser"
      passwordFile: "/etc/prompower-auth/password"
      caFile: "/etc/prompower-ca/ca.crt"
      nodeDeviceMapFile: ""
      query: "device_power_watts_avg{room=\"R3.033\",rack=\"Rack 3\",device=\"${POWER_DEVICE}\"}"

vmagent:
  enabled: true
  scrape:
    interval: 5s
    target: "127.0.0.1:28282"
    path: /metrics
    scheme: http
  remoteWrite:
    url: "https://prometheus.cs.hm.edu:8428/api/v1/write"
    username: "vmuser"
    passwordSecretName: vm-remote-write-auth
    passwordSecretKey: password
    caFile: "/etc/prompower-ca/ca.crt"
```

### Worker-Only Scheduling

For production, schedule Kepler only on worker nodes (not control-plane nodes) by using nodeAffinity/tolerations in the DaemonSet. This avoids interference with control-plane components and keeps resource usage predictable.

### Production Rollout Checklist

- [x] custom image built and pushed to GitLab registry (`gitlab.lrz.de:5005/green-it/students/bachelor/eder-moritz/kepler-prompower:v0.1.2`)
- [x] deploy via Helm chart at `manifests/helm/kepler` with `manifests/helm/kepler/values-prompower-prod.yaml`
- [x] monitoring endpoint reachable from cluster
- [x] HTTPS and Basic Auth working for central VictoriaMetrics endpoint
- [x] secrets and (optional) ConfigMap present in namespace
- [x] production node names match external device labels OR optional node-to-device mapping is configured
- [x] worker-only scheduling configured for DaemonSet
- [x] vmagent `remote_write` forwarding verified and external power input query returns values

### Smoke Test After Rollout

1. Verify DaemonSet and pods (example assumes release name `kepler`):

```bash
kubectl get daemonset kepler -n default
kubectl get pods -n default -l app.kubernetes.io/name=kepler
```

If your release name is different, replace `kepler` with your release name.

2. Port-forward a Kepler pod and test metrics endpoint:

```bash
kubectl port-forward -n default svc/kepler 28282:28282
curl http://localhost:28282/metrics | grep kepler_node_cpu_watts
```


3. Query central VictoriaMetrics for forwarded metrics:

```bash
curl -G "https://prometheus.cs.hm.edu:8428/api/v1/query" --data-urlencode 'query=sum by (node_name)(kepler_node_cpu_watts)' --user "vmuser:yourpassword" --cacert /etc/prompower-ca/ca.crt
```

### Notes for Cluster Migration

When migrating this setup to a different cluster you typically need to adjust:

- image registry and tag
- central monitoring endpoint hostname
- CA certificate secret
- auth secrets
- optional node-to-device mapping ConfigMap (only if node names differ from external device labels)
- scheduling constraints (node selectors/taints)
- dashboards and queries with environment-specific labels

Portable example queries:

```promql
sum(kepler_node_cpu_watts)
```

```promql
topk(10, sum by (node_name, comm, pid) (kepler_process_cpu_watts))
```

## �📖 Documentation

### User Documentation

- **[Installation Guide](docs/user/installation.md)** - Detailed installation instructions for all deployment methods
- **[Configuration Guide](docs/user/configuration.md)** - Configuration options and examples
- **[Metrics Documentation](docs/user/metrics.md)** - Available metrics and their descriptions

### Developer Documentation

- **[Architecture Documentation](docs/developer/design/architecture/)** - Complete architectural documentation including design principles, system components, data flow, concurrency model, and deployment patterns
- **[Power Attribution Guide](docs/developer/power-attribution-guide.md)** - How Kepler measures and attributes power consumption
- **[Developer Documentation](docs/developer/)** - Contributing guidelines and development workflow

For more detailed documentation, please visit the [official Kepler documentation](https://sustainable-computing.io/kepler/).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For more detailed information about contributing to this project, please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) file.

### Gen AI policy

Our project adheres to the Linux Foundation's Generative AI Policy, which can be viewed at [https://www.linuxfoundation.org/legal/generative-ai](https://www.linuxfoundation.org/legal/generative-ai).

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=sustainable-computing-io/kepler&type=Date)](https://www.star-history.com/#sustainable-computing-io/kepler&Date)

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSES](LICENSES) for details.
