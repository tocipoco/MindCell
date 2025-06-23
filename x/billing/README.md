# Billing Module

## Overview

Handles pay-per-inference fee calculation and multi-party revenue distribution.

## Revenue Split

- Model Owner: 60%
- Shard Nodes: 30%
- Protocol Treasury: 10%

## Fee Calculation

`Total Fee = Base Fee + (Compute Units × Compute Price)`

## Messages

- `MsgProcessPayment`: Process inference payment
- `MsgSettleBilling`: Finalize billing record
- `MsgUpdateFeeConfig`: Update fee parameters

## Configuration

- BaseFee: 1000 mcell
- ComputeUnitPrice: 10 mcell
- Revenue percentages (governance adjustable)
