package model

type StageTemplate struct {
	BaseModel
	Name           string `gorm:"column:name"`
	Type           string `gorm:"column:type"`
	DockerImage    string `gorm:"column:docker_image"`
	DockerImageTag string `gorm:"column:docker_image_tag"` // e.g., 3.9.9, latest
}

func (StageTemplate) TableName() string {
	return "stage_templates"
}
