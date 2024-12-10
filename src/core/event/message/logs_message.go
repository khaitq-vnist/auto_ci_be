package message

type LogsEventMessage struct {
	UserId      int64 `json:"user_id"`
	ProjectId   int64 `json:"project_id"`
	PipelineId  int64 `json:"pipeline_id"`
	ExecutionId int64 `json:"execution_id"`
}
