package model

type ProjectModel struct {
	BaseModel
	Name        string  `gorm:"size:255;not null;column:name"`
	FullName    string  `gorm:"size:255;not null;column:full_name"`
	Private     bool    `gorm:"default:false;column:private"`
	OwnerID     uint    `gorm:"not null;column:owner_id"`
	Description *string `gorm:"type:text;column:description"`
	Language    *string `gorm:"size:50;column:language"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
