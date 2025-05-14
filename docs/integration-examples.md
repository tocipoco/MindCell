# Integration Examples

## Web Application Integration

### React.js Example

```javascript
import { MindCellClient } from '@mindcell/sdk';
import { useState } from 'react';

function ImageClassifier() {
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);
  
  const client = new MindCellClient({
    endpoint: process.env.REACT_APP_MINDCELL_RPC
  });

  const classifyImage = async (imageFile) => {
    setLoading(true);
    try {
      const base64 = await fileToBase64(imageFile);
      const inference = await client.inference.submit({
        modelId: 42,
        inputData: { image: base64 },
        maxFee: '500mcell'
      });
      
      const result = await client.inference.waitForResult(inference.requestId);
      setResult(result.output);
    } catch (error) {
      console.error('Inference failed:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <input type="file" onChange={(e) => classifyImage(e.target.files[0])} />
      {loading && <p>Processing...</p>}
      {result && <p>Predicted: {result.class} ({result.confidence}%)</p>}
    </div>
  );
}
```

## Backend Integration

### Node.js Express API

```javascript
const express = require('express');
const { MindCellClient } = require('@mindcell/sdk');

const app = express();
const client = new MindCellClient({
  endpoint: process.env.MINDCELL_RPC,
  mnemonic: process.env.MINDCELL_MNEMONIC
});

app.post('/api/inference', async (req, res) => {
  try {
    const { modelId, inputData } = req.body;
    
    const result = await client.inference.submit({
      modelId,
      inputData,
      maxFee: '1000mcell'
    });
    
    const inference = await client.inference.waitForResult(result.requestId);
    
    res.json({
      success: true,
      output: inference.result,
      fee: inference.fee
    });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

app.listen(3000);
```

### Python Flask API

```python
from flask import Flask, request, jsonify
from mindcell import MindCellClient
import os

app = Flask(__name__)
client = MindCellClient(
    endpoint=os.getenv('MINDCELL_RPC'),
    mnemonic=os.getenv('MINDCELL_MNEMONIC')
)

@app.route('/api/inference', methods=['POST'])
def inference():
    data = request.json
    
    try:
        result = client.inference.submit(
            model_id=data['model_id'],
            input_data=data['input_data'],
            max_fee='1000mcell'
        )
        
        inference = client.inference.wait_for_result(result.request_id)
        
        return jsonify({
            'success': True,
            'output': inference.result,
            'fee': inference.fee
        })
    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    app.run(port=5000)
```

## Mobile Integration

### iOS Swift Example

```swift
import MindCellSDK

class InferenceService {
    let client: MindCellClient
    
    init() {
        client = MindCellClient(
            endpoint: "https://rpc.mindcell.network",
            chainId: "mindcell-1"
        )
    }
    
    func submitInference(modelId: UInt64, image: UIImage) async throws -> InferenceResult {
        let base64 = image.jpegData(compressionQuality: 0.8)?.base64EncodedString()
        
        let request = InferenceRequest(
            modelId: modelId,
            inputData: ["image": base64],
            maxFee: "500mcell"
        )
        
        let result = try await client.inference.submit(request)
        return try await client.inference.waitForResult(result.requestId)
    }
}
```

### Android Kotlin Example

```kotlin
import com.mindcell.sdk.MindCellClient
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

class InferenceRepository {
    private val client = MindCellClient(
        endpoint = "https://rpc.mindcell.network",
        chainId = "mindcell-1"
    )

    suspend fun classifyImage(modelId: Long, imageBase64: String): InferenceResult {
        return withContext(Dispatchers.IO) {
            val request = InferenceRequest(
                modelId = modelId,
                inputData = mapOf("image" to imageBase64),
                maxFee = "500mcell"
            )
            
            val submission = client.inference.submit(request)
            client.inference.waitForResult(submission.requestId)
        }
    }
}
```

## Batch Processing

### Batch Inference Service

```python
from mindcell import MindCellClient
from concurrent.futures import ThreadPoolExecutor
import asyncio

class BatchInferenceService:
    def __init__(self):
        self.client = MindCellClient('https://rpc.mindcell.network')
        self.client.load_wallet(os.getenv('MNEMONIC'))
    
    async def process_batch(self, model_id, inputs, batch_size=32):
        results = []
        
        for i in range(0, len(inputs), batch_size):
            batch = inputs[i:i+batch_size]
            
            # Submit batch request
            batch_result = self.client.inference.submit_batch(
                model_id=model_id,
                inputs=batch,
                max_fee=f'{len(batch) * 500}mcell'
            )
            
            # Wait for results
            batch_outputs = await self.client.inference.wait_for_batch(
                batch_result.batch_id
            )
            
            results.extend(batch_outputs)
        
        return results

# Usage
service = BatchInferenceService()
inputs = [{'image': img1}, {'image': img2}, ...]
results = asyncio.run(service.process_batch(42, inputs))
```

## Streaming Integration

### Real-time Text Generation

```javascript
const stream = client.inference.submitStreaming({
  modelId: 789,
  inputData: { prompt: 'Write a technical blog post about zkML' },
  maxFee: '3000mcell'
});

let fullText = '';

stream.on('token', (token) => {
  fullText += token;
  updateUI(fullText); // Update UI in real-time
});

stream.on('complete', (result) => {
  console.log('Generation complete');
  console.log('Total cost:', result.actualFee);
  saveToDB(fullText);
});

stream.on('error', (error) => {
  console.error('Stream error:', error);
  showErrorToUser(error.message);
});
```

## Monitoring Integration

### Prometheus Metrics Collection

```yaml
# prometheus.yml
scrape_configs:
  - job_name: 'mindcell'
    static_configs:
      - targets: ['localhost:26660']
    metrics_path: '/metrics'
    scrape_interval: 15s
```

### Grafana Dashboard Query

```javascript
// Inference requests per minute
rate(mindcell_inference_requests_total[1m])

// Average inference latency
avg(mindcell_inference_duration_seconds)

// Node reputation scores
mindcell_node_reputation_score
```

More integration examples available at:
https://github.com/mindcell-network/integration-examples
