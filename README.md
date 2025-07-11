# Nyx

Nyx is a GoLang demo application designed for validating Kubernetes controllers, Prometheus metrics, and modern cloud-native best practices. It provides a simple HTTP API and exposes Prometheus metrics, making it ideal for testing, learning, and integration with Kubernetes monitoring stacks.

---

## Features

- **HTTP API**: Simple endpoints for demo and health checks.
- **Prometheus Metrics**: Exposes HTTP request counters and histograms for monitoring.
- **Kubernetes Native**: Includes manifests for deployment, service, and ServiceMonitor.
- **Configurable**: Supports configuration via YAML file, environment variables, and command-line flags.
- **Graceful Shutdown**: Handles SIGINT/SIGTERM for clean exits.

---

## Getting Started

### Prerequisites

- Go 1.24+
- Docker (optional, for container builds)
- Kubernetes cluster (for deployment)
- [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator) (for ServiceMonitor support)

---

### Running Locally

1. **Clone the repository**
    ```sh
    git clone https://github.com/MatteoMori/nyx.git
    cd nyx
    ```

2. **Build and run**
    ```sh
    go build -o nyx
    ./nyx start
    ```

3. **Access the API**
    - HTTP server: [http://localhost:8080/hello](http://localhost:8080/hello)
    - Metrics: [http://localhost:9090/metrics](http://localhost:9090/metrics)

---

### Configuration

Nyx supports configuration via:
- YAML file (`nyx.yaml`)
- Environment variables
- Command-line flags

**Example `nyx.yaml`:**
```yaml
prometheusPort: "9090"
verbosity: 1
```

**Environment variable override:**
```sh
export PROMETHEUSPORT=12345
./nyx start
```

---

### Kubernetes Deployment

1. **Apply the manifests**
    ```sh
    kubectl apply -f manifests/
    ```

2. **Expose metrics to Prometheus**
    ```sh
    kubectl apply -f manifests/servicemonitor.yaml
    ```

3. **Access the app**
    - HTTP: via the `nyx` service on port 80
    - Metrics: via the `nyx-metrics` service on port 9090

---

## Metrics

Nyx exposes the following Prometheus metrics:

- `http_requests_total{path,method,status}`: Total number of HTTP requests.
- `http_request_duration_seconds{path,method}`: Histogram of request durations.

Metrics are available at `/metrics` on the configured Prometheus port (default: 9090).

---

## Project Structure

```
.
├── cmd/                # CLI entrypoints (Cobra)
├── manifests/          # Kubernetes manifests
├── pkg/
│   ├── nyx/            # Application logic and metrics
│   └── shared/         # Shared config types
├── main.go             # Main entrypoint
├── nyx.yaml            # Example config
├── Dockerfile          # Multi-stage build for container images
└── README.md           # This file
```
