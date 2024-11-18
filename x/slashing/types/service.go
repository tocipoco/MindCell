package types

import "context"

type MsgServer interface {
	SlashNode(context.Context, *MsgSlashNode) (*MsgSlashNodeResponse, error)
	UpdateSlashingParams(context.Context, *MsgUpdateSlashingParams) (*MsgUpdateSlashingParamsResponse, error)
}

type MsgSlashNodeResponse struct {
	Success bool `json:"success"`
}

type MsgUpdateSlashingParamsResponse struct {
	Success bool `json:"success"`
}

type QueryServer interface {
	GetSlashingRecord(context.Context, *QueryGetSlashingRecordRequest) (*QueryGetSlashingRecordResponse, error)
	GetSlashingParams(context.Context, *QueryGetSlashingParamsRequest) (*QueryGetSlashingParamsResponse, error)
	ListSlashingRecords(context.Context, *QueryListSlashingRecordsRequest) (*QueryListSlashingRecordsResponse, error)
}

type QueryGetSlashingRecordRequest struct {
	RecordID uint64 `json:"record_id"`
}

type QueryGetSlashingRecordResponse struct {
	Record SlashingRecord `json:"record"`
}

type QueryGetSlashingParamsRequest struct{}

type QueryGetSlashingParamsResponse struct {
	Params SlashingParams `json:"params"`
}

type QueryListSlashingRecordsRequest struct {
	NodeAddress string `json:"node_address,omitempty"`
}

type QueryListSlashingRecordsResponse struct {
	Records []SlashingRecord `json:"records"`
}

func RegisterMsgServer(server interface{}, impl MsgServer)           {}
func RegisterQueryServer(server interface{}, impl QueryServer)       {}
func RegisterQueryHandlerClient(ctx context.Context, mux, client interface{}) error { return nil }
func NewQueryClient(ctx interface{}) QueryClient                     { return nil }

type QueryClient interface{}
var _Msg_serviceDesc = struct{}{}

