package model

import "time"

type PipelineHistoryModel struct {
	BaseModel
	PipelineID  uint       `gorm:"index;not null;foreignKey:PipelineID;constraint:OnDelete:CASCADE;column:pipeline_id"`
	StartedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP;column:started_at"`
	EndedAt     *time.Time `gorm:"column:ended_at"`
	Status      string     `gorm:"size:50;not null;column:status"`
	TriggeredBy uint       `gorm:"not null;column:triggered_by"`
	Logs        *string    `gorm:"type:text;column:logs"`
}
