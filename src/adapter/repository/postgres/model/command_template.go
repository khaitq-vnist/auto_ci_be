package model

type CommandTemplateModel struct {
	BaseModel
	Command string `gorm:"column:command"`
	StageId int64  `gorm:"column:stage_id"`
}

func (CommandTemplateModel) TableName() string {
	return "commands_templates"
}
