# zkML Integration Guide

## Overview

MindCell uses zero-knowledge machine learning (zkML) proofs to ensure inference correctness without revealing the full model.

## Proof Generation

Shard nodes generate partial proofs for their computation:

```go
proof := generatePartialProof(input, modelShard)
```

## Proof Aggregation

The InferenceGateway aggregates partial proofs:

```go
finalProof := aggregateProofs(partialProofs)
```

## Verification

Proofs are verified on-chain before billing:

```go
valid := verifyProof(finalProof, publicInputs)
```

## Supported Proof Systems

- gnark
- Halo2
- Custom SNARK implementations

## Performance Considerations

- Proof generation: ~100-500ms per shard
- Verification: ~10-50ms on-chain
- Proof size: 200-500 bytes

