package model

type ProviderModel struct {
	BaseModel
	Name    string `gorm:"column:name"`
	Code    string `gorm:"column:code"`
	ApiUrl  string `gorm:"column:api_url"`
	AuthUrl string `gorm:"column:auth_url"`
}

func (ProviderModel) TableName() string {
	return "providers"
}
