package model

type StageCommandTemplateModel struct {
	BaseModel
	StageID   int64 `gorm:"column:stage_id"`
	CommandID int64 `gorm:"column:command_id"`
}

func (StageCommandTemplateModel) TableName() string {
	return "stage_command_templates"
}
