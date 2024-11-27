package entity

type PipelineTemplateEntity struct {
	BaseEntity
	Name        string
	BuildTool   string
	Description string
}
