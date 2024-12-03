package response

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type BuddyEventResponse struct {
	Type string   `json:"type"`
	Refs []string `json:"refs"`
}

type BuddyPipelineResponse struct {
	ID     int64                 `json:"id"`
	Name   string                `json:"name"`
	On     string                `json:"on"`
	Refs   []string              `json:"refs"`
	Events []*BuddyEventResponse `json:"events"`
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
