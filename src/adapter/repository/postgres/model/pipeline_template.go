package model

type PipelineTemplate struct {
	BaseModel
	Name        string `gorm:"column:name"`
	BuildTool   string `gorm:"column:build_tool"`
	Description string `gorm:"column:description"`
}

func (PipelineTemplate) TableName() string {
	return "pipeline_template"
}
