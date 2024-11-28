package model

type PipelineTemplateModel struct {
	BaseModel
	Name        string `gorm:"column:name"`
	BuildTool   string `gorm:"column:build_tool"`
	Description string `gorm:"column:description"`
}

func (PipelineTemplateModel) TableName() string {
	return "pipeline_templates"
}
