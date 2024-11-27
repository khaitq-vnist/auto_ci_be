package entity

type StageTemplateEntity struct {
	BaseEntity
	Name           string
	Type           string
	DockerImage    string
	DockerImageTag string
}
