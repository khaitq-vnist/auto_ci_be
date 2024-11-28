package entity

type CommandTemplateEntity struct {
	BaseEntity
	Command string
	StageID int64
}
