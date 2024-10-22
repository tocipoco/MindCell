package types

import (
	context "context"
)

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterNode(context.Context, *MsgRegisterNode) (*MsgRegisterNodeResponse, error)
	AssignShard(context.Context, *MsgAssignShard) (*MsgAssignShardResponse, error)
	ReplaceShard(context.Context, *MsgReplaceShard) (*MsgReplaceShardResponse, error)
	UpdateNodeReputation(context.Context, *MsgUpdateNodeReputation) (*MsgUpdateNodeReputationResponse, error)
}

// Response types
type MsgRegisterNodeResponse struct {
	Success bool `json:"success"`
}

type MsgAssignShardResponse struct {
	Success bool `json:"success"`
}

type MsgReplaceShardResponse struct {
	Success bool `json:"success"`
}

type MsgUpdateNodeReputationResponse struct {
	Success bool `json:"success"`
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetShardAssignment(context.Context, *QueryGetShardAssignmentRequest) (*QueryGetShardAssignmentResponse, error)
	GetNodeInfo(context.Context, *QueryGetNodeInfoRequest) (*QueryGetNodeInfoResponse, error)
	ListNodes(context.Context, *QueryListNodesRequest) (*QueryListNodesResponse, error)
}

// Query request/response types
type QueryGetShardAssignmentRequest struct {
	ModelID uint64 `json:"model_id"`
	ShardID uint32 `json:"shard_id"`
}

type QueryGetShardAssignmentResponse struct {
	Assignment ShardAssignment `json:"assignment"`
}

type QueryGetNodeInfoRequest struct {
	NodeAddress string `json:"node_address"`
}

type QueryGetNodeInfoResponse struct {
	Node       NodeInfo       `json:"node"`
	Reputation NodeReputation `json:"reputation"`
}

type QueryListNodesRequest struct {
	ActiveOnly bool `json:"active_only"`
}

type QueryListNodesResponse struct {
	Nodes []NodeInfo `json:"nodes"`
}

// RegisterMsgServer registers the msg server
func RegisterMsgServer(server interface{}, impl MsgServer) {}

// RegisterQueryServer registers the query server
func RegisterQueryServer(server interface{}, impl QueryServer) {}

// RegisterQueryHandlerClient registers the query handler client
func RegisterQueryHandlerClient(ctx context.Context, mux interface{}, client interface{}) error {
	return nil
}

// NewQueryClient creates a new query client
func NewQueryClient(ctx interface{}) QueryClient {
	return nil
}

// QueryClient is the client API for Query service.
type QueryClient interface {
	GetShardAssignment(ctx context.Context, in *QueryGetShardAssignmentRequest) (*QueryGetShardAssignmentResponse, error)
	GetNodeInfo(ctx context.Context, in *QueryGetNodeInfoRequest) (*QueryGetNodeInfoResponse, error)
	ListNodes(ctx context.Context, in *QueryListNodesRequest) (*QueryListNodesResponse, error)
}

var _Msg_serviceDesc = struct{}{}

