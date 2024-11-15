package entity

type PipelineHistoryEntity struct {
	BaseEntity
	PipelineId  int64
	StartedAt   int64
	EndedAt     *int64
	Status      string
	TriggeredBy uint
	Logs        *string
}
