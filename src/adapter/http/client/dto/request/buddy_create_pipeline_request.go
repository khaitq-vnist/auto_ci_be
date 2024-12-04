package request

import (
	"encoding/json"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type BuddyCreatePipelineRequest struct {
	Name                      string           `json:"name,omitempty"`
	ON                        string           `json:"on,omitempty"`
	Refs                      []string         `json:"refs,omitempty"`
	Event                     []*BuddyEventReq `json:"events,omitempty"`
	ExecutionMessageTemplate  string           `json:"execution_message_template,omitempty"`
	AlwaysFromScratch         bool             `json:"always_from_scratch,omitempty"`
	AutoClearCache            bool             `json:"auto_clear_cache,omitempty"`
	NoSkipToMostRecent        bool             `json:"no_skip_to_most_recent,omitempty"`
	DoNotCreateCommitStatus   bool             `json:"do_not_create_commit_status,omitempty"`
	IgnoreFailOnProjectStatus bool             `json:"ignore_fail_on_project_status,omitempty"`
	FilesystemChangesetBase   string           `json:"filesystem_changeset_base,omitempty"`
	ConcurrentPipelineRuns    bool             `json:"concurrent_pipeline_runs,omitempty"`
	GitChangesetBase          string           `json:"git_changeset_base,omitempty"`
}

type BuddyEventReq struct {
	Type string   `json:"type"`
	Refs []string `json:"refs"`
}

func ToBuddyPipelineRequest(pipeline *entity.PipelineEntity) *BuddyCreatePipelineRequest {
	events := make([]*BuddyEventReq, 0)
	for _, event := range pipeline.Events {
		events = append(events, &BuddyEventReq{
			Type: event.Type,
			Refs: event.Refs,
		})
	}
	return &BuddyCreatePipelineRequest{
		Name:  pipeline.Name,
		ON:    pipeline.On,
		Event: events,
		Refs:  pipeline.Refs,
	}
}
func (b *BuddyCreatePipelineRequest) ToJson() ([]byte, error) {
	return json.Marshal(b)
}
