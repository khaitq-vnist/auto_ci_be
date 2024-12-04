package entity

type PipelineEntity struct {
	ID                  int64
	Name                string
	On                  string
	Refs                []string
	Events              []*EventEntity
	Actions             []*ActionEntity
	LastExecutionStatus string
	LastExecuteRevision string
}
type EventEntity struct {
	Type string
	Refs []string
}
