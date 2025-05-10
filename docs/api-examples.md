# API Usage Examples

## Model Registration

### Example 1: Register Image Classification Model

```bash
# Upload model metadata to IPFS
ipfs add model-metadata.json
# Returns: QmXYZ123...

# Register model
mindcelld tx modelregistry register-model \
  --metadata-cid="QmXYZ123..." \
  --shard-count=4 \
  --from=model-owner \
  --fees=1000mcell
```

### Example 2: Update Model Version

```bash
# Upload updated metadata
ipfs add model-metadata-v2.json

# Update existing model
mindcelld tx modelregistry update-model \
  --model-id=42 \
  --metadata-cid="QmABC456..." \
  --from=model-owner
```

## Inference Requests

### Example 1: Image Classification

```javascript
const client = new MindCellClient({endpoint: 'https://rpc.mindcell.network'});

const result = await client.inference.submit({
  modelId: 42,
  inputData: {
    image: base64EncodedImage,
    preprocessing: 'resize_224x224'
  },
  maxFee: '500mcell'
});

console.log('Predicted class:', result.output.class);
console.log('Confidence:', result.output.confidence);
```

### Example 2: Text Generation

```python
from mindcell import MindCellClient

client = MindCellClient('https://rpc.mindcell.network')
client.load_wallet(mnemonic)

result = client.inference.submit(
    model_id=123,
    input_data={
        'prompt': 'Explain quantum computing',
        'max_tokens': 500,
        'temperature': 0.7
    },
    max_fee='2000mcell'
)

print(result.output['text'])
```

## Node Operations

### Example 1: Register as Shard Node

```bash
mindcelld tx shardallocator register-node \
  --endpoint="https://node.example.com:8080" \
  --stake=50000mcell \
  --max-shards=100 \
  --from=node-operator
```

### Example 2: Claim Rewards

```bash
# Check pending rewards
mindcelld query reward node-reward cosmos1abc...

# Claim rewards
mindcelld tx reward claim-reward \
  --from=node-operator
```

## More Examples

Find complete code examples at:
https://github.com/mindcell-network/examples
