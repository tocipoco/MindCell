# ModelRegistry Module

## Overview

The ModelRegistry module manages the registration, versioning, and metadata of AI models on the MindCell network.

## Key Features

- Model registration with IPFS metadata
- Version tracking and updates
- Owner-based access control
- Model activation/deactivation
- Query by owner or ID

## State

- Models: Map of model ID to Model struct
- ModelsCount: Total number of registered models
- Owner Index: Mapping from owner address to model IDs

## Messages

- `MsgRegisterModel`: Register a new AI model
- `MsgUpdateModel`: Update existing model metadata
- `MsgDeactivateModel`: Deactivate a model

## Queries

- `GetModel`: Retrieve model by ID
- `ListModels`: List all models (optionally filtered by owner)
- `ModelsCount`: Get total count of registered models

## Events

- `model_registered`: Emitted when new model is registered
- `model_updated`: Emitted when model is updated
- `model_deactivated`: Emitted when model is deactivated
