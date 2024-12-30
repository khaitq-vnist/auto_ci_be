package model

type ProjectModel struct {
	BaseModel
	Name                string  `gorm:"size:255;not null;column:name"`
	FullName            string  `gorm:"size:255;not null;column:full_name"`
	Private             bool    `gorm:"default:false;column:private"`
	OwnerID             uint    `gorm:"not null;column:owner_id"`
	Description         *string `gorm:"type:text;column:description"`
	Language            *string `gorm:"size:50;column:language"`
	HtmlUrl             string  `gorm:"size:255;column:html_url"`
	ProviderRepoID      int64   `gorm:"column:provider_repo_id"`
	IntegrationId       int64   `gorm:"column:integration_id"`
	ThirdPartyProjectID string  `gorm:"column:third_party_project_id"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}
