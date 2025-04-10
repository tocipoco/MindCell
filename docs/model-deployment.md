# Model Deployment Guide

## Prerequisites

Before deploying your AI model to MindCell:

- Model trained and validated
- Model exported to ONNX or compatible format
- IPFS node access for metadata storage
- MCELL tokens for registration fee
- MindCell wallet with sufficient balance

## Step 1: Model Preparation

### Export Your Model

**PyTorch to ONNX:**
```python
import torch
import torch.onnx

model = YourModel()
model.load_state_dict(torch.load('model.pth'))
model.eval()

dummy_input = torch.randn(1, 3, 224, 224)
torch.onnx.export(
    model,
    dummy_input,
    "model.onnx",
    input_names=['input'],
    output_names=['output'],
    dynamic_axes={'input': {0: 'batch_size'}}
)
```

**TensorFlow to ONNX:**
```python
import tf2onnx
import tensorflow as tf

spec = (tf.TensorSpec((None, 224, 224, 3), tf.float32, name="input"),)
model_proto, _ = tf2onnx.convert.from_keras(model, input_signature=spec)

with open("model.onnx", "wb") as f:
    f.write(model_proto.SerializeToString())
```

### Quantization (Optional)

Reduce model size and improve inference speed:

```python
# PyTorch quantization
quantized_model = torch.quantization.quantize_dynamic(
    model, {torch.nn.Linear}, dtype=torch.qint8
)
```

**Benefits:**
- 4x smaller model size
- 2-4x faster inference
- Lower hosting costs

## Step 2: Model Sharding

### Determine Shard Count

```python
model_size_mb = os.path.getsize('model.onnx') / (1024 * 1024)

# Recommended: 10-50MB per shard
target_shard_size = 25  # MB
shard_count = max(3, int(model_size_mb / target_shard_size))

print(f"Recommended shard count: {shard_count}")
```

### Shard the Model

```python
import numpy as np

def shard_model(model_path, num_shards):
    with open(model_path, 'rb') as f:
        model_bytes = f.read()
    
    shard_size = len(model_bytes) // num_shards
    shards = []
    
    for i in range(num_shards):
        start = i * shard_size
        end = start + shard_size if i < num_shards - 1 else len(model_bytes)
        shard = model_bytes[start:end]
        
        # Encrypt shard
        encrypted_shard = encrypt_aes256(shard, get_shard_key(i))
        shards.append(encrypted_shard)
        
        # Upload to temporary storage
        shard_cid = upload_to_ipfs(encrypted_shard)
        print(f"Shard {i} uploaded: {shard_cid}")
    
    return shards

shards = shard_model('model.onnx', shard_count)
```

## Step 3: Metadata Preparation

Create comprehensive model metadata:

```json
{
  "name": "Image Classification Model",
  "version": "1.0.0",
  "description": "ResNet50 fine-tuned on custom dataset",
  "architecture": "ResNet50",
  "framework": "PyTorch",
  "format": "ONNX",
  "parameters": 25000000,
  "size_mb": 98,
  "shard_count": 4,
  "input_shape": [1, 3, 224, 224],
  "output_shape": [1, 1000],
  "inference_time_ms": 45,
  "accuracy": {
    "top1": 0.924,
    "top5": 0.987
  },
  "dataset": "ImageNet + Custom",
  "license": "Apache-2.0",
  "tags": ["computer-vision", "classification", "resnet"],
  "shards": [
    "QmShard1CIDxxx...",
    "QmShard2CIDxxx...",
    "QmShard3CIDxxx...",
    "QmShard4CIDxxx..."
  ],
  "checksum": "sha256:abc123..."
}
```

### Upload Metadata to IPFS

```bash
# Upload metadata JSON
ipfs add metadata.json

# Result: QmMetadataCIDxxx...
```

## Step 4: Register on MindCell

### Using CLI

```bash
mindcelld tx modelregistry register-model \
  --metadata-cid="QmMetadataCIDxxx..." \
  --shard-count=4 \
  --from=mykey \
  --chain-id=mindcell-1 \
  --fees=1000mcell \
  --gas=auto
```

### Using Go SDK

```go
import (
    "github.com/tocipoco/MindCell/x/modelregistry/types"
)

msg := types.NewMsgRegisterModel(
    ownerAddress,
    "QmMetadataCIDxxx...",
    4, // shard count
)

txResponse, err := client.BroadcastTx(msg)
if err != nil {
    log.Fatal(err)
}

modelID := extractModelID(txResponse)
fmt.Printf("Model registered with ID: %d\n", modelID)
```

### Using JavaScript SDK

```javascript
const { MindCellClient } = require('@mindcell/sdk');

const client = new MindCellClient({
  endpoint: 'https://rpc.mindcell.network',
  chainId: 'mindcell-1'
});

const result = await client.modelRegistry.register({
  metadataCid: 'QmMetadataCIDxxx...',
  shardCount: 4,
  signerAddress: walletAddress
});

console.log('Model ID:', result.modelId);
```

## Step 5: Monitor Deployment

### Check Model Status

```bash
# Query model information
mindcelld query modelregistry get-model <model-id>

# Check shard assignments
mindcelld query shardallocator list-shards --model-id=<model-id>
```

### Verify Shard Distribution

```bash
# Check which nodes host your shards
for shard_id in {0..3}; do
  mindcelld query shardallocator get-assignment \
    --model-id=<model-id> \
    --shard-id=$shard_id
done
```

## Step 6: Test Inference

### Submit Test Request

```bash
# Prepare test input
echo '{"data": [0.5, 0.3, ...]}' > input.json

# Submit inference request
mindcelld tx inferencegateway submit-inference \
  --model-id=<model-id> \
  --input-file=input.json \
  --fee=500mcell \
  --from=mykey
```

### Monitor Request Status

```bash
# Check request status
mindcelld query inferencegateway get-request <request-id>

# Wait for completion
while true; do
  status=$(mindcelld query inferencegateway get-request <request-id> | jq -r '.status')
  if [ "$status" = "completed" ]; then
    echo "Inference complete!"
    break
  fi
  sleep 2
done
```

## Step 7: Production Deployment

### Set Pricing

Update your model's inference pricing:

```bash
# Set price per inference
mindcelld tx modelregistry update-pricing \
  --model-id=<model-id> \
  --base-price=100mcell \
  --compute-multiplier=1.5 \
  --from=mykey
```

### Enable Auto-scaling

Configure automatic shard replication based on demand:

```bash
mindcelld tx shardallocator configure-autoscale \
  --model-id=<model-id> \
  --min-replicas=3 \
  --max-replicas=10 \
  --target-utilization=0.70 \
  --from=mykey
```

### Set Up Monitoring

```bash
# Subscribe to model events
mindcelld events subscribe \
  --query="model_id=$MODEL_ID" \
  --output=json \
  | tee model-events.log
```

## Best Practices

### Security

1. **Encrypt Shards**: Use AES-256 for shard encryption
2. **Key Management**: Store decryption keys in HSM or secure vault
3. **Access Control**: Limit who can update model metadata
4. **Audit Trail**: Monitor all model access and modifications

### Reliability

1. **Redundancy**: Deploy at least 3 replica shards
2. **Health Checks**: Implement liveness probes
3. **Failover**: Configure automatic shard replacement
4. **Backups**: Keep model backups on multiple storage providers

### Cost Optimization

1. **Right-size Shards**: Balance between cost and performance
2. **Compression**: Use model compression techniques
3. **Pruning**: Remove unnecessary parameters
4. **Caching**: Cache frequently used results

### Compliance

1. **Data Privacy**: Ensure training data compliance (GDPR, etc.)
2. **Model Cards**: Document model capabilities and limitations
3. **Bias Testing**: Evaluate for fairness and bias
4. **Version Control**: Maintain model versioning history

## Updating Models

### Deploy New Version

```bash
# Upload new metadata
ipfs add metadata_v2.json

# Update model
mindcelld tx modelregistry update-model \
  --model-id=<model-id> \
  --metadata-cid="QmNewMetadataCID..." \
  --from=mykey
```

### Gradual Rollout

1. Deploy new version as separate model
2. Route 10% of traffic to new version
3. Monitor performance and errors
4. Gradually increase traffic
5. Deprecate old version after validation

## Troubleshooting

### Shard Assignment Failures

**Issue**: Not enough nodes accepting shards

**Solution:**
- Increase registration fee to attract more nodes
- Lower shard size requirements
- Contact node operators directly

### High Inference Latency

**Issue**: Slow inference responses

**Solution:**
- Check shard node locations (geographic distribution)
- Optimize model (quantization, pruning)
- Increase number of shard replicas
- Use edge nodes closer to users

### Failed Proofs

**Issue**: zkML proof verification failures

**Solution:**
- Verify model format compatibility
- Check shard integrity
- Update to latest proof system
- Report bug if persistent

## Support

- Documentation: https://docs.mindcell.network
- Discord: https://discord.gg/mindcell
- Email: support@mindcell.network
- GitHub Issues: Report bugs and feature requests
