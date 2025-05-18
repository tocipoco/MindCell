# Monitoring and Observability Guide

## Metrics Collection

### Prometheus Setup

Install and configure Prometheus to scrape MindCell metrics:

```yaml
# prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'mindcell-validator'
    static_configs:
      - targets: ['localhost:26660']
        labels:
          instance: 'validator-1'
          env: 'production'
  
  - job_name: 'mindcell-node'
    static_configs:
      - targets: ['node1:26660', 'node2:26660']
```

### Key Metrics

**Consensus Metrics:**
- `tendermint_consensus_height` - Current block height
- `tendermint_consensus_validators` - Active validator count
- `tendermint_consensus_missing_validators` - Validators not signing
- `tendermint_consensus_byzantine_validators` - Byzantine behavior detected

**Network Metrics:**
- `tendermint_p2p_peers` - Connected peer count
- `tendermint_p2p_peer_receive_bytes_total` - Data received
- `tendermint_p2p_peer_send_bytes_total` - Data sent

**Application Metrics:**
- `mindcell_inference_requests_total` - Total inference requests
- `mindcell_inference_duration_seconds` - Request processing time
- `mindcell_model_count` - Registered models
- `mindcell_active_nodes` - Active shard nodes
- `mindcell_node_reputation_score` - Node reputation scores

## Grafana Dashboards

### Import Pre-built Dashboard

```bash
# Download dashboard JSON
curl -O https://raw.githubusercontent.com/mindcell-network/monitoring/main/grafana/mindcell-overview.json

# Import in Grafana UI: Dashboards > Import > Upload JSON
```

### Custom Dashboard Queries

**Inference Rate:**
```promql
rate(mindcell_inference_requests_total[5m])
```

**Success Rate:**
```promql
rate(mindcell_inference_success_total[5m]) / 
rate(mindcell_inference_requests_total[5m])
```

**P95 Latency:**
```promql
histogram_quantile(0.95, 
  rate(mindcell_inference_duration_seconds_bucket[5m])
)
```

## Logging

### Structured Logging Configuration

```toml
# app.toml
log_level = "info"
log_format = "json"
```

### Log Aggregation (ELK Stack)

```yaml
# filebeat.yml
filebeat.inputs:
  - type: log
    enabled: true
    paths:
      - /var/log/mindcell/*.log
    json.keys_under_root: true
    json.add_error_key: true

output.elasticsearch:
  hosts: ["localhost:9200"]
  index: "mindcell-logs-%{+yyyy.MM.dd}"
```

### Important Log Patterns

Monitor for these patterns:
- `"level":"error"` - Application errors
- `"module":"consensus".*"missing"` - Consensus issues
- `"slashed"` - Slashing events
- `"proof.*failed"` - Proof verification failures

## Alerting

### Alert Rules

```yaml
# alert_rules.yml
groups:
  - name: mindcell_critical
    rules:
      - alert: NodeDown
        expr: up{job="mindcell-validator"} == 0
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "MindCell node is down"
          description: "Node {{ $labels.instance }} has been down for 5 minutes"

      - alert: HighInferenceLatency
        expr: histogram_quantile(0.95, rate(mindcell_inference_duration_seconds_bucket[5m])) > 5
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "High inference latency detected"

      - alert: LowNodeReputation
        expr: mindcell_node_reputation_score < 50
        for: 1h
        labels:
          severity: warning
        annotations:
          summary: "Node reputation below threshold"
```

### AlertManager Configuration

```yaml
# alertmanager.yml
route:
  group_by: ['alertname', 'instance']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 12h
  receiver: 'team-notifications'

receivers:
  - name: 'team-notifications'
    slack_configs:
      - api_url: 'YOUR_SLACK_WEBHOOK_URL'
        channel: '#mindcell-alerts'
        title: 'MindCell Alert'
        text: '{{ range .Alerts }}{{ .Annotations.summary }}{{ end }}'
```

## Tracing

### Jaeger Integration

```go
import "github.com/opentracing/opentracing-go"

// In keeper methods
func (k Keeper) ProcessInference(ctx sdk.Context, req InferenceRequest) error {
    span, ctx := opentracing.StartSpanFromContext(ctx, "ProcessInference")
    defer span.Finish()
    
    span.SetTag("model_id", req.ModelID)
    span.SetTag("requester", req.Requester)
    
    // Processing logic
}
```

## Health Checks

### Liveness Probe

```bash
#!/bin/bash
# health-check.sh
STATUS=$(mindcelld status 2>&1)
if echo "$STATUS" | jq -e '.SyncInfo.catching_up == false' > /dev/null; then
    exit 0
else
    exit 1
fi
```

### Readiness Probe

```bash
# Check if node is synced and has peers
PEERS=$(curl -s localhost:26657/net_info | jq '.result.n_peers | tonumber')
SYNCED=$(mindcelld status 2>&1 | jq -r '.SyncInfo.catching_up')

if [ "$PEERS" -gt 0 ] && [ "$SYNCED" = "false" ]; then
    exit 0
else
    exit 1
fi
```

## Performance Monitoring

### Resource Usage

```bash
# CPU and Memory
pidstat -p $(pgrep mindcelld) 5

# Disk I/O
iostat -x 5

# Network
iftop -i eth0
```

### Database Performance

```bash
# LevelDB stats
du -sh ~/.mindcell/data

# Query performance logging
# Enable in config.toml:
# log_level = "debug"
```

## Dashboards

### Node Operator Dashboard

Panels to include:
- Block height and sync status
- Peer connections
- Validator voting power
- Missed blocks counter
- Resource utilization (CPU, RAM, Disk)
- Network bandwidth

### Model Owner Dashboard

Panels to include:
- Total inference requests
- Request success rate
- Average latency
- Revenue earned
- Active vs inactive models
- User engagement metrics

## Best Practices

1. **Set up alerts before problems occur**
2. **Monitor trends, not just absolute values**
3. **Create runbooks for common issues**
4. **Review metrics weekly**
5. **Test alerting system regularly**
6. **Document baseline performance**
7. **Track SLOs (Service Level Objectives)**

## Tools

- **Metrics**: Prometheus + Grafana
- **Logging**: ELK Stack or Loki
- **Tracing**: Jaeger or Zipkin
- **APM**: DataDog or New Relic
- **Uptime**: UptimeRobot or Pingdom
