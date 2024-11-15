package model

type Pipeline struct {
	BaseModel
	Name        string                  `gorm:"size:255;not null;column:name"`
	Description *string                 `gorm:"type:text;column:description"`
	OwnerID     uint                    `gorm:"not null;column:owner_id"`
	IsActive    bool                    `gorm:"default:true;column:is_active"`
	ProjectID   uint64                  `gorm:"index;foreignKey:ProjectID;constraint:OnDelete:SET NULL;column:project_id"`
	Stages      []PipelineStageModel    `gorm:"foreignKey:PipelineID;column:stages"`
	Variables   []PipelineVariableModel `gorm:"foreignKey:PipelineID;column:variables"`
	Services    []PipelineServiceModel  `gorm:"foreignKey:PipelineID;column:services"`
	History     []PipelineHistoryModel  `gorm:"foreignKey:PipelineID;column:history"`
}
