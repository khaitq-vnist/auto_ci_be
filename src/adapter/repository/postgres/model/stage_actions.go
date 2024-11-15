package model

type StageActionModel struct {
	BaseModel
	StageID          uint    `gorm:"index;not null;foreignKey:StageID;constraint:OnDelete:CASCADE;column:stage_id"`
	Name             string  `gorm:"size:255;not null;column:name"`
	ActionOrder      int     `gorm:"not null;column:action_order"`
	Command          string  `gorm:"type:text;not null;column:command"`
	DockerImageName  *string `gorm:"size:255;column:docker_image_name"`
	DockerImageTag   *string `gorm:"size:50;column:docker_image_tag"`
	WorkingDirectory *string `gorm:"size:255;column:working_directory"`
	ShellType        *string `gorm:"size:50;default:'bash';column:shell_type"`
}
