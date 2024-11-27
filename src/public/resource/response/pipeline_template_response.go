package response

// PipelineTemplateResponse is a response struct for pipeline template
type PipelineTemplateResponse struct {
	ID          int64                    `json:"id"`
	Name        string                   `json:"name"`
	BuildTool   string                   `json:"build_tool"`
	Description string                   `json:"description"`
	Stages      []*StageTemplateResponse `json:"stages"`
}

type StageTemplateResponse struct {
	ID             int64                      `json:"id"`
	Name           string                     `json:"name"`
	Type           string                     `json:"type"`
	DockerImage    string                     `json:"docker_image"`
	DockerImageTag string                     `json:"docker_image_tag"`
	Commands       []*CommandTemplateResponse `json:"commands"`
	Variables      []*VariableResponse        `json:"variables"`
}

type CommandTemplateResponse struct {
	ID      int64  `json:"id"`
	Command string `json:"command"`
}
type VariableResponse struct {
	ID    int64  `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PipelineSettingsResponse struct {
	ID        int64                   `json:"id"`
	Status    string                  `json:"status"`
	Trigger   PipelineTriggerResponse `json:"trigger"`
	CodeScope CodeScopeResponse       `json:"code_scope"`
}
type PipelineTriggerResponse struct {
	ID          int64  `json:"id"`
	TriggerType string `json:"trigger_type"`
	TriggerName string `json:"trigger_name"`
}

type CodeScopeResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
