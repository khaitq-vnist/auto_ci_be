package model

type PipelineStageTemplateModel struct {
	BaseModel
	TemplateID int64 `gorm:"column:template_id"`
	StageID    int64 `gorm:"column:stage_id"`
}

func (PipelineStageTemplateModel) TableName() string {
	return "pipeline_stage_templates"
}
