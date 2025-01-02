package model

type ExecutionHistoryModel struct {
	BaseModel
	ProjectID         int64   `gorm:"column:project_id"`
	PipelineID        int64   `gorm:"column:pipeline_id"`
	ThirdPartyID      int64   `gorm:"column:third_party_id"`
	ThirdPartyProject string  `gorm:"column:third_party_project"`
	LogsFile          string  `gorm:"column:logs_file"`
	Coverage          float64 `gorm:"column:coverage"`
}

func (ExecutionHistoryModel) TableName() string {
	return "execution_history"
}
