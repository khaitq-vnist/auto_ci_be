package entity

type ProjectEntity struct {
	BaseEntity
	Name        string
	FullName    string
	Private     bool
	OwnerID     uint
	Description *string
	Language    *string
}
