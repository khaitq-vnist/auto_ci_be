package request

import "github.com/khaitq-vnist/auto_ci_be/core/entity"

type BuddyCreatePipelineRequest struct {
	Name                      string           `json:"name"`
	ON                        string           `json:"on"`
	Refs                      []string         `json:"refs"`
	Event                     []*BuddyEventReq `json:"events"`
	ExecutionMessageTemplate  string           `json:"execution_message_template"`
	AlwaysFromScratch         bool             `json:"always_from_scratch"`
	AutoClearCache            bool             `json:"auto_clear_cache"`
	NoSkipToMostRecent        bool             `json:"no_skip_to_most_recent"`
	DoNotCreateCommitStatus   bool             `json:"do_not_create_commit_status"`
	IgnoreFailOnProjectStatus bool             `json:"ignore_fail_on_project_status"`
	FilesystemChangesetBase   string           `json:"filesystem_changeset_base"`
	ConcurrentPipelineRuns    bool             `json:"concurrent_pipeline_runs"`
	GitChangesetBase          string           `json:"git_changeset_base"`
}
type BuddyEventReq struct {
	Type string `json:"type"`
	refs []string
}

func ToBuddyPipelineRequest(pipeline *entity.PipelineEntity) *BuddyCreatePipelineRequest {
	events := make([]*BuddyEventReq, 0)
	for _, event := range pipeline.Events {
		events = append(events, &BuddyEventReq{
			Type: event.Type,
			refs: event.Refs,
		})
	}
	return &BuddyCreatePipelineRequest{
		Name:  pipeline.Name,
		ON:    pipeline.On,
		Event: events,
		Refs:  pipeline.Refs,
	}
}
