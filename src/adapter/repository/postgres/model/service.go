package model

import "gorm.io/datatypes"

type ServiceModel struct {
	BaseModel
	Type       string         `gorm:"column:type"`
	Version    string         `gorm:"column:version"`
	Connection datatypes.JSON `gorm:"column:connection"`
}

func (ServiceModel) TableName() string {
	return "services"
}
