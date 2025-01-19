# MindCell API Documentation

## Model Registry API

### Register Model
```
POST /mindcell/modelregistry/register
```

Register a new AI model to the network.

### Get Model
```
GET /mindcell/modelregistry/models/{model_id}
```

Retrieve model information by ID.

### List Models
```
GET /mindcell/modelregistry/models
```

List all registered models.

## Shard Allocator API

### Register Node
```
POST /mindcell/shardallocator/nodes/register
```

Register as a shard hosting node.

### Assign Shard
```
POST /mindcell/shardallocator/shards/assign
```

Assign a shard to a node.

## Inference Gateway API

### Submit Inference
```
POST /mindcell/inference/submit
```

Submit an inference request.

### Get Request Status
```
GET /mindcell/inference/requests/{request_id}
```

Check the status of an inference request.

## Billing API

### Get Billing Record
```
GET /mindcell/billing/records/{request_id}
```

Retrieve billing information for a request.

### Get Fee Configuration
```
GET /mindcell/billing/config
```

Get current fee configuration.

