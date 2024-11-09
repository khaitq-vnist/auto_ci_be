package model

type IntegrationModel struct {
	BaseModel
	Name             string `gorm:"column:name"`
	UserId           int64  `gorm:"column:user_id"`
	ProviderId       int64  `gorm:"column:provider_id"`
	ProviderName     string `gorm:"column:provider_name"`
	AccessToken      string `gorm:"column:access_token"`
	ProviderUsername string `gorm:"column:provider_username"`
}

func (IntegrationModel) TableName() string {
	return "integrations"
}
