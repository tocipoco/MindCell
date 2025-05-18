# Testing Guide

## Test Structure

MindCell uses multiple testing layers:

### Unit Tests

Test individual components in isolation:

```go
func TestKeeperGetModel(t *testing.T) {
    keeper, ctx := setupTestKeeper(t)
    
    // Setup
    expectedModel := types.Model{
        ID: 1,
        Owner: "cosmos1test",
        ShardCount: 4,
    }
    keeper.SetModel(ctx, expectedModel)
    
    // Test
    actualModel, found := keeper.GetModel(ctx, 1)
    
    // Assert
    require.True(t, found)
    require.Equal(t, expectedModel.ID, actualModel.ID)
    require.Equal(t, expectedModel.Owner, actualModel.Owner)
}
```

### Integration Tests

Test module interactions:

```go
func TestInferenceFlow(t *testing.T) {
    // Setup test network
    network := testutil.NewNetwork(t, 3)
    defer network.Cleanup()
    
    // Register model
    modelID := registerTestModel(t, network)
    
    // Assign shards
    assignShards(t, network, modelID, 4)
    
    // Submit inference
    requestID := submitInference(t, network, modelID)
    
    // Verify result
    result := getInferenceResult(t, network, requestID)
    require.NotEmpty(t, result.Output)
}
```

### End-to-End Tests

Test complete user workflows:

```bash
#!/bin/bash
# e2e_test.sh

# Start local testnet
./scripts/init-testnet.sh

# Register model
mindcelld tx modelregistry register-model \
  --metadata-cid="QmTest123" \
  --shard-count=2 \
  --from=test1 \
  --yes

# Wait for confirmation
sleep 6

# Query model
MODEL_ID=$(mindcelld query modelregistry list-models --output=json | jq '.models[0].id')

# Submit inference
mindcelld tx inferencegateway submit-inference \
  --model-id=$MODEL_ID \
  --input='{"data": [1,2,3]}' \
  --fee=500mcell \
  --from=test1 \
  --yes

# Verify result
sleep 6
mindcelld query inferencegateway list-requests
```

## Running Tests

### All Tests

```bash
make test
```

### Specific Module

```bash
go test ./x/modelregistry/... -v
```

### With Coverage

```bash
make test-coverage
```

### Race Detection

```bash
go test -race ./...
```

### Benchmarks

```bash
go test -bench=. -benchmem ./x/shardallocator/keeper
```

## Writing Good Tests

### Test Naming

```go
// Good: Descriptive test names
func TestKeeper_GetModel_ReturnsModelWhenExists(t *testing.T)
func TestKeeper_GetModel_ReturnsFalseWhenNotFound(t *testing.T)

// Bad: Unclear names
func TestGetModel1(t *testing.T)
func TestGM(t *testing.T)
```

### Table-Driven Tests

```go
func TestValidateMessage(t *testing.T) {
    tests := []struct {
        name      string
        msg       types.MsgRegisterModel
        expectErr bool
        errMsg    string
    }{
        {
            name: "valid message",
            msg: types.MsgRegisterModel{
                Owner: "cosmos1test",
                MetadataCID: "QmValid123",
                ShardCount: 4,
            },
            expectErr: false,
        },
        {
            name: "empty metadata CID",
            msg: types.MsgRegisterModel{
                Owner: "cosmos1test",
                MetadataCID: "",
                ShardCount: 4,
            },
            expectErr: true,
            errMsg: "metadata CID cannot be empty",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.msg.ValidateBasic()
            if tt.expectErr {
                require.Error(t, err)
                require.Contains(t, err.Error(), tt.errMsg)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

### Test Fixtures

```go
// testdata/fixtures.go
package testdata

var (
    ValidModel = types.Model{
        ID: 1,
        Owner: "cosmos1test",
        MetadataCID: "QmTest123",
        ShardCount: 4,
        Active: true,
    }
    
    ValidInferenceRequest = types.InferenceRequest{
        ModelID: 1,
        Requester: "cosmos1requester",
        InputData: `{"data": [1,2,3]}`,
    }
)
```

## Mocking

### Interface Mocks

```go
type MockShardAllocator struct {
    SelectBestNodeFunc func(ctx sdk.Context) (string, error)
}

func (m *MockShardAllocator) SelectBestNode(ctx sdk.Context) (string, error) {
    if m.SelectBestNodeFunc != nil {
        return m.SelectBestNodeFunc(ctx)
    }
    return "cosmos1mocknode", nil
}

// Usage in tests
func TestWithMock(t *testing.T) {
    mock := &MockShardAllocator{
        SelectBestNodeFunc: func(ctx sdk.Context) (string, error) {
            return "cosmos1specific", nil
        },
    }
    
    // Test using mock
}
```

## CI/CD Testing

### GitHub Actions Workflow

Automatically runs on every PR:
- Unit tests
- Integration tests  
- Linting
- Coverage reporting
- Build verification

### Pre-commit Hooks

```bash
# .git/hooks/pre-commit
#!/bin/bash
make test
make lint
```

## Test Coverage Goals

- Overall: >80%
- Critical paths: 100%
- New code: >90%

Check coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Performance Testing

### Load Testing

```bash
# Use hey for load testing
hey -n 10000 -c 100 \
  -m POST \
  -H "Content-Type: application/json" \
  -d '{"model_id": 42, "input": {"data": [1,2,3]}}' \
  http://localhost:1317/mindcell/inference/submit
```

### Benchmark Tests

```go
func BenchmarkProcessInference(b *testing.B) {
    keeper, ctx := setupBenchKeeper(b)
    req := getTestRequest()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        keeper.ProcessInference(ctx, req)
    }
}
```

## Best Practices

1. Test one thing per test
2. Use descriptive test names
3. Arrange-Act-Assert pattern
4. Clean up resources in tests
5. Use test helpers for common setup
6. Mock external dependencies
7. Test edge cases and error conditions
8. Keep tests fast and independent
9. Use parallel tests when possible
10. Document complex test scenarios
