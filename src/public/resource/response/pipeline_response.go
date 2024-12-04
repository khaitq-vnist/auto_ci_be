package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type PipelineResponse struct {
	ID                  int64    `json:"id"`
	Name                string   `json:"name"`
	On                  string   `json:"on"`
	Refs                []string `json:"refs"`
	LastExecutionStatus string   `json:"last_execution_status"`
	LastExecuteRevision string   `json:"last_execute_revision"`
}

func ToListPipelineResponse(pipelineEntity []*entity.PipelineEntity) []*PipelineResponse {
	pipelines := make([]*PipelineResponse, 0)
	for _, pipeline := range pipelineEntity {
		pipelines = append(pipelines, &PipelineResponse{
			ID:                  pipeline.ID,
			Name:                pipeline.Name,
			On:                  pipeline.On,
			Refs:                pipeline.Refs,
			LastExecutionStatus: pipeline.LastExecutionStatus,
			LastExecuteRevision: pipeline.LastExecuteRevision,
		})
	}
	return pipelines
}
