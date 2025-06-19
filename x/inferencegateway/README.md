# InferenceGateway Module

## Overview

Routes inference requests, manages zkML proof verification, and tracks request lifecycle.

## Key Features

- Inference request submission
- zkML proof verification
- Request status tracking
- Result delivery
- Fee estimation

## State

- InferenceRequests: All submitted requests
- RequestCount: Total requests processed
- Verification results

## Messages

- `MsgSubmitInference`: Submit new inference request
- `MsgVerifyProof`: Verify zkML proof
- `MsgCompleteInference`: Mark inference as complete

## Request Lifecycle

1. Submitted (pending)
2. Processing (shards computing)
3. Proof verification
4. Completed or Failed

## Proof Systems

- gnark
- Halo2
- Custom SNARK implementations
