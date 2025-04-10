# Client SDK Guide

## Installation

### JavaScript/TypeScript

```bash
npm install @mindcell/sdk
# or
yarn add @mindcell/sdk
```

### Python

```bash
pip install mindcell-sdk
```

### Go

```bash
go get github.com/tocipoco/MindCell/sdk
```

## Quick Start

### JavaScript Example

```javascript
const { MindCellClient } = require('@mindcell/sdk');

// Initialize client
const client = new MindCellClient({
  endpoint: 'https://rpc.mindcell.network',
  chainId: 'mindcell-1',
  prefix: 'cosmos'
});

// Connect wallet
await client.connectWallet(mnemonic);

// Submit inference request
const result = await client.inference.submit({
  modelId: 123,
  inputData: { image: base64Image },
  maxFee: '1000mcell'
});

console.log('Request ID:', result.requestId);

// Wait for result
const inference = await client.inference.waitForResult(result.requestId);
console.log('Output:', inference.result);
```

### Python Example

```python
from mindcell import MindCellClient

# Initialize client
client = MindCellClient(
    endpoint='https://rpc.mindcell.network',
    chain_id='mindcell-1'
)

# Load wallet
client.load_wallet(mnemonic='your mnemonic here')

# Submit inference
result = client.inference.submit(
    model_id=123,
    input_data={'image': base64_image},
    max_fee='1000mcell'
)

# Async wait for completion
inference = await client.inference.wait_for_result(result.request_id)
print(f"Output: {inference.result}")
```

### Go Example

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/tocipoco/MindCell/sdk"
)

func main() {
    client := sdk.NewClient(sdk.Config{
        Endpoint: "https://rpc.mindcell.network",
        ChainID:  "mindcell-1",
    })
    
    // Load wallet
    wallet, _ := sdk.LoadWallet(mnemonic)
    client.SetWallet(wallet)
    
    // Submit inference
    req := &sdk.InferenceRequest{
        ModelID:   123,
        InputData: inputMap,
        MaxFee:    "1000mcell",
    }
    
    result, err := client.SubmitInference(context.Background(), req)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Request ID: %d\n", result.RequestID)
    
    // Poll for result
    inference, _ := client.WaitForInference(context.Background(), result.RequestID)
    fmt.Printf("Output: %v\n", inference.Result)
}
```

## API Reference

### Model Registry

#### List Available Models

```javascript
const models = await client.modelRegistry.listModels({
  owner: 'cosmos1...', // optional filter
  active: true
});

models.forEach(model => {
  console.log(`${model.id}: ${model.metadata.name}`);
});
```

#### Get Model Details

```python
model = client.model_registry.get_model(model_id=123)
print(f"Owner: {model.owner}")
print(f"Shards: {model.shard_count}")
print(f"Version: {model.version}")
```

### Inference Gateway

#### Submit Inference Request

```typescript
interface InferenceRequest {
  modelId: number;
  inputData: any;
  maxFee: string;
  timeout?: number; // seconds
}

const result = await client.inference.submit({
  modelId: 123,
  inputData: {
    prompt: "Translate to French: Hello world"
  },
  maxFee: '500mcell',
  timeout: 30
});
```

#### Query Request Status

```python
status = client.inference.get_status(request_id=456)

print(f"Status: {status.status}")  # pending, processing, completed, failed
print(f"Progress: {status.progress}%")
print(f"Estimated time: {status.eta_seconds}s")
```

#### Stream Results (WebSocket)

```javascript
client.inference.streamResults(requestId, (update) => {
  console.log('Progress:', update.progress);
  console.log('Partial result:', update.partialResult);
});
```

### Node Operations

#### Register as Node Operator

```python
response = client.shard_allocator.register_node(
    endpoint='https://my-node.example.com',
    stake_amount='10000mcell',
    max_shards=50
)

print(f"Node registered: {response.success}")
```

#### Claim Rewards

```javascript
const rewards = await client.rewards.getPending(nodeAddress);
console.log('Pending rewards:', rewards.amount);

if (rewards.amount > '0') {
  const tx = await client.rewards.claim();
  console.log('Claimed:', tx.amount);
}
```

### Billing

#### Get Fee Estimate

```python
estimate = client.billing.estimate_fee(
    model_id=123,
    compute_units=1000
)

print(f"Estimated cost: {estimate.total_fee}")
print(f"Breakdown:")
print(f"  Base fee: {estimate.base_fee}")
print(f"  Compute: {estimate.compute_fee}")
```

#### Query Billing History

```javascript
const history = await client.billing.getHistory({
  requester: myAddress,
  startDate: '2025-01-01',
  endDate: '2025-12-31'
});

const totalSpent = history.reduce((sum, record) => 
  sum + parseFloat(record.totalFee), 0
);

console.log('Total spent:', totalSpent, 'MCELL');
```

## Advanced Usage

### Batch Inference

Process multiple inputs efficiently:

```python
requests = [
    {'image': img1_base64},
    {'image': img2_base64},
    {'image': img3_base64}
]

# Submit batch (10x cheaper than individual)
batch_result = client.inference.submit_batch(
    model_id=123,
    inputs=requests,
    max_fee='2000mcell'
)

# Wait for all results
results = await client.inference.wait_for_batch(batch_result.batch_id)
for i, result in enumerate(results):
    print(f"Result {i}: {result.output}")
```

### Streaming Inference

For long-running inferences (e.g., LLM text generation):

```javascript
const stream = client.inference.submitStreaming({
  modelId: 789,
  inputData: { prompt: 'Write a story about...' },
  maxFee: '2000mcell'
});

stream.on('token', (token) => {
  process.stdout.write(token);
});

stream.on('complete', (result) => {
  console.log('\n\nInference complete');
  console.log('Total cost:', result.actualFee);
});
```

### Error Handling

```python
from mindcell.exceptions import (
    InsufficientFundsError,
    ModelNotFoundError,
    ProofVerificationError
)

try:
    result = client.inference.submit(
        model_id=123,
        input_data=data
    )
except InsufficientFundsError as e:
    print(f"Need more MCELL tokens: {e.required_amount}")
except ModelNotFoundError:
    print("Model does not exist or is inactive")
except ProofVerificationError:
    print("Inference proof validation failed - please retry")
```

### Custom Timeout and Retry

```javascript
const options = {
  timeout: 60000,  // 60 seconds
  retries: 3,
  retryDelay: 1000,
  onRetry: (attempt) => {
    console.log(`Retry attempt ${attempt}...`);
  }
};

const result = await client.inference.submit(request, options);
```

## Event Subscription

### Subscribe to Events

```python
# Subscribe to model events
def on_model_event(event):
    if event.type == 'model_registered':
        print(f"New model: {event.attributes['model_id']}")
    elif event.type == 'model_updated':
        print(f"Model updated: {event.attributes['version']}")

client.events.subscribe('modelregistry', on_model_event)

# Subscribe to your inference requests
def on_inference_update(event):
    print(f"Status: {event.attributes['status']}")
    if event.attributes['status'] == 'completed':
        print(f"Result: {event.attributes['result']}")

client.events.subscribe(
    'inferencegateway',
    on_inference_update,
    filter={'requester': my_address}
)
```

### WebSocket Connection

```javascript
const ws = client.connectWebSocket();

ws.on('open', () => {
  ws.subscribe({
    query: `inference.requester='${myAddress}'`
  });
});

ws.on('event', (event) => {
  console.log('Event received:', event);
});
```

## Configuration

### Client Options

```typescript
interface ClientConfig {
  endpoint: string;           // RPC endpoint
  chainId: string;           // Chain identifier
  prefix?: string;           // Address prefix (default: 'cosmos')
  gasPrice?: string;         // Gas price (default: '0.025mcell')
  gasAdjustment?: number;    // Gas adjustment factor (default: 1.5)
  timeout?: number;          // Request timeout ms (default: 30000)
  broadcastMode?: string;    // 'sync' | 'async' | 'block'
}
```

### Environment Variables

```bash
export MINDCELL_RPC_ENDPOINT="https://rpc.mindcell.network"
export MINDCELL_CHAIN_ID="mindcell-1"
export MINDCELL_GAS_PRICE="0.025mcell"
export MINDCELL_MNEMONIC="your twenty four word mnemonic..."
```

## Testing

### Mock Client for Testing

```python
from mindcell.testing import MockClient

# Use mock client in tests
client = MockClient()
client.set_mock_response('inference.submit', {
    'request_id': 999,
    'status': 'completed',
    'result': {'class': 'cat', 'confidence': 0.95}
})

# Your code will receive mock data
result = client.inference.submit(...)
assert result.request_id == 999
```

## Performance Tips

1. **Connection Pooling**: Reuse client instances
2. **Batch Requests**: Group multiple inferences
3. **Async Operations**: Use async/await patterns
4. **Caching**: Cache model metadata locally
5. **Compression**: Compress large input data

## Troubleshooting

### Common Issues

**Connection Timeout:**
```python
client = MindCellClient(
    endpoint='https://rpc.mindcell.network',
    timeout=60000  # Increase timeout
)
```

**Insufficient Gas:**
```javascript
const result = await client.inference.submit(request, {
  gas: 'auto',
  gasAdjustment: 2.0  // Increase gas limit
});
```

**Invalid Address Format:**
```python
# Ensure correct address prefix
from mindcell.utils import validate_address

if not validate_address(address, prefix='cosmos'):
    raise ValueError("Invalid address format")
```

## Examples Repository

Find complete working examples at:
https://github.com/mindcell-network/sdk-examples

- Image classification
- Text generation
- Object detection
- Sentiment analysis
- And more...
