package model

type PipelineServiceModel struct {
	BaseModel
	PipelineID    uint   `gorm:"index;not null;foreignKey:PipelineID;constraint:OnDelete:CASCADE;column:pipeline_id"`
	ServiceType   string `gorm:"size:50;not null;column:service_type"`
	Version       string `gorm:"size:50;not null;column:version"`
	Configuration string `gorm:"type:jsonb;column:configuration"`
}
