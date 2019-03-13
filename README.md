# This is a Simple logger to demo the use of Promtail as a sidecar to send logs to Loki

## Prerequisites

This docs assume:
 - you have loki and grafana already deployed. Please refered to official documentation for installation
 - The logfile you want to scrape is in JSON format

This Helm chart deploy a application pod with 2 containers:
    - a Golang app writing logs in a separate file.
    - a Promtail that read that log file and send it to loki.

The file path can be updated via the [./helm/values.yaml](./helm/values.yaml) file.

`sidecar.labels` is a map where you can add the labels that will be added to your log entry in Loki.

Example:

- `Logfile` located at `/home/slog/creator.log`
- Adding labels
    - `job: promtail-sidecar`
    - `test: golang`
```yaml
sidecar:
  logfile:
    path: /home/slog
    filename: creator.log
  labels:
    job: promtail-sidecar
    test: golang
```

