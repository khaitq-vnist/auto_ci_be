package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type BuddyEventResponse struct {
	Type string   `json:"type"`
	Refs []string `json:"refs"`
}

type BuddyPipelineResponse struct {
	ID                  int64                 `json:"id"`
	Name                string                `json:"name"`
	On                  string                `json:"on"`
	Refs                []string              `json:"refs"`
	Events              []*BuddyEventResponse `json:"events"`
	LastExecutionStatus string                `json:"last_execution_status"`
	LastExecuteRevision string                `json:"last_execute_revision"`
}

type BuddyPipelineListResponse struct {
	Pipelines []*BuddyPipelineResponse `json:"pipelines"`
}

func ToListPipelineEntities(buddyRsp *BuddyPipelineListResponse) []*entity.PipelineEntity {
	pipelines := make([]*entity.PipelineEntity, 0)
	for _, pipeline := range buddyRsp.Pipelines {
		pipelines = append(pipelines, &entity.PipelineEntity{
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
func ToPipelineEntity(buddyRsp *BuddyPipelineResponse) *entity.PipelineEntity {
	events := make([]*entity.EventEntity, 0)
	for _, event := range buddyRsp.Events {
		events = append(events, &entity.EventEntity{
			Type: event.Type,
			Refs: event.Refs,
		})
	}
	return &entity.PipelineEntity{
		ID:     buddyRsp.ID,
		Name:   buddyRsp.Name,
		On:     buddyRsp.On,
		Refs:   buddyRsp.Refs,
		Events: events,
	}
}
