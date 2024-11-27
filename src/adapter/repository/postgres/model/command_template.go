package model

type CommandTemplateModel struct {
	BaseModel
	Command string `gorm:"column:command"`
}

func (CommandTemplateModel) TableName() string {
	return "command_templates"
}
