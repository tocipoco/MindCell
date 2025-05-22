# Data Structures Reference

## Core Types

### Model

Represents a registered AI model in the ModelRegistry:

```go
type Model struct {
    ID          uint64  // Unique identifier
    Owner       string  // Cosmos address of model owner
    MetadataCID string  // IPFS CID of model metadata
    ShardCount  uint32  // Number of shards
    Version     uint32  // Current version number
    Active      bool    // Whether model is active
    CreatedAt   int64   // Unix timestamp of registration
    UpdatedAt   int64   // Unix timestamp of last update
}
```

### ShardAssignment

Maps model shards to nodes:

```go
type ShardAssignment struct {
    ModelID     uint64  // Model identifier
    ShardID     uint32  // Shard number (0-indexed)
    NodeAddress string  // Cosmos address of hosting node
    AssignedAt  int64   // Timestamp of assignment
    Status      string  // active, inactive, slashed
}
```

### InferenceRequest

Tracks inference requests:

```go
type InferenceRequest struct {
    RequestID   uint64  // Unique request identifier
    ModelID     uint64  // Target model
    Requester   string  // Address that submitted request
    InputData   string  // JSON-encoded input data
    Status      string  // pending, processing, completed, failed
    Result      string  // JSON-encoded output
    ProofHash   string  // Hash of zkML proof
    SubmittedAt int64   // Submission timestamp
    CompletedAt int64   // Completion timestamp
    Fee         string  // Total fee paid
}
```

### NodeInfo

Information about shard nodes:

```go
type NodeInfo struct {
    Address          string  // Node's Cosmos address
    Endpoint         string  // HTTP/gRPC endpoint
    StakeAmount      string  // Staked MCELL amount
    MaxShards        uint32  // Maximum shard capacity
    CurrentShards    uint32  // Currently hosting
    Active           bool    // Operational status
    RegistrationTime int64   // When registered
}
```

### NodeReputation

Tracks node performance:

```go
type NodeReputation struct {
    NodeAddress      string   // Node identifier
    ReputationScore  float64  // 0-100 score
    TotalInferences  uint64   // All-time inference count
    SuccessfulCount  uint64   // Successful inferences
    FailedCount      uint64   // Failed inferences
    LastActivityTime int64    // Most recent activity
}
```

## Storage Keys

### ModelRegistry

- `0x01 | modelID` → Model
- `0x02` → ModelsCount
- `0x03 | owner | modelID` → ModelID (index)

### ShardAllocator

- `0x01 | modelID | shardID` → ShardAssignment
- `0x02 | nodeAddress` → NodeReputation
- `0x03 | nodeAddress | shardID` → ShardID (index)
- `0x04 | nodeAddress` → NodeInfo

### InferenceGateway

- `0x01 | requestID` → InferenceRequest
- `0x02 | proofHash` → ValidationResult
- `0x03 | requester | requestID` → RequestID (index)

## Events

### ModelRegistry Events

```go
sdk.NewEvent(
    "model_registered",
    sdk.NewAttribute("model_id", "123"),
    sdk.NewAttribute("owner", "cosmos1..."),
    sdk.NewAttribute("shard_count", "4"),
)
```

### InferenceGateway Events

```go
sdk.NewEvent(
    "inference_submitted",
    sdk.NewAttribute("request_id", "456"),
    sdk.NewAttribute("model_id", "123"),
    sdk.NewAttribute("requester", "cosmos1..."),
)
```

## Message Types

### MsgRegisterModel

```go
type MsgRegisterModel struct {
    Owner       string  // Model owner address
    MetadataCID string  // IPFS metadata CID
    ShardCount  uint32  // Number of shards
}
```

### MsgSubmitInference

```go
type MsgSubmitInference struct {
    Requester string  // Request submitter
    ModelID   uint64  // Target model
    InputData string  // JSON input
    Fee       string  // Payment amount
}
```

## State Transitions

### Model Registration Flow

```
Initial State: ModelsCount = N

1. Receive MsgRegisterModel
2. Validate message
3. Create Model{ID: N+1, ...}
4. Store Model
5. Increment ModelsCount
6. Create owner index
7. Emit event

Final State: ModelsCount = N+1, Model exists
```

### Inference Flow

```
1. MsgSubmitInference received
2. Create InferenceRequest{Status: "pending"}
3. Store request
4. Emit event
5. [Off-chain: Shards compute, generate proofs]
6. MsgCompleteInference received  
7. Verify proof
8. Update request{Status: "completed"}
9. Process billing
10. Distribute rewards
```

## Genesis State

```go
type GenesisState struct {
    ModelRegistry    ModelRegistryGenesis
    ShardAllocator   ShardAllocatorGenesis
    InferenceGateway InferenceGatewayGenesis
    Billing          BillingGenesis
    Reward           RewardGenesis
    Slashing         SlashingGenesis
    Token            TokenGenesis
}
```

Each module exports/imports its own genesis state for network initialization or export.
