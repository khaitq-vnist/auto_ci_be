package entity

type ExecutionHistoryEntity struct {
	BaseEntity
	ProjectID         int64
	PipelineID        int64
	ThirdPartyID      int64
	ThirdPartyProject string
	LogsFile          string
	Coverage          float64
}
