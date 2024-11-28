package entity

type VariableTemplateEntity struct {
	BaseEntity
	StageID int64
	Key     string
	Value   string
}
