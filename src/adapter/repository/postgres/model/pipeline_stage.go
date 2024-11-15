package model

type PipelineStageModel struct {
	BaseModel
	PipelineID uint   `gorm:"index;not null;foreignKey:PipelineID;constraint:OnDelete:CASCADE;column:pipeline_id"`
	Name       string `gorm:"size:255;not null;column:name"`
	StageOrder int    `gorm:"not null;column:stage_order"`
	Type       string `gorm:"size:50;not null;column:type"`

	Actions []StageActionModel `gorm:"foreignKey:StageID;column:actions"`
}
