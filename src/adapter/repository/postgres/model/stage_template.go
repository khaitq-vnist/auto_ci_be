package model

type StageTemplateModel struct {
	BaseModel
	PipelineTemplateID int64  `gorm:"column:pipeline_id"`
	Name               string `gorm:"column:name"`
	Type               string `gorm:"column:type"`
	DockerImage        string `gorm:"column:docker_image"`
	DockerImageTag     string `gorm:"column:docker_image_tag"` // e.g., 3.9.9, latest
}

func (StageTemplateModel) TableName() string {
	return "stage_templates"
}
