package entity

type StageTemplateEntity struct {
	BaseEntity
	Name               string
	Type               string
	DockerImage        string
	DockerImageTag     string
	PipelineTemplateID int64
	Commands           []*CommandTemplateEntity
	Variables          []*VariableTemplateEntity
}
