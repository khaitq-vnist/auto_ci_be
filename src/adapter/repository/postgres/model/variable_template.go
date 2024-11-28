package model

type VariableTemplateModel struct {
	BaseModel
	StageID int64  `gorm:"column:stage_id"`
	Key     string `gorm:"column:key"`
	Value   string `gorm:"column:value"`
}

func (VariableTemplateModel) TableName() string {
	return "variable_templates"
}
