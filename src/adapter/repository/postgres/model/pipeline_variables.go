package model

type PipelineVariableModel struct {
	BaseModel
	PipelineID uint   `gorm:"index;not null;foreignKey:PipelineID;constraint:OnDelete:CASCADE;column:pipeline_id"`
	Key        string `gorm:"size:255;not null;column:key"`
	Value      string `gorm:"type:text;not null;column:value"`
	IsSecret   bool   `gorm:"default:false;column:is_secret"`
}
