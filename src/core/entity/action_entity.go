package entity

type ActionEntity struct {
	ID              int64
	Name            string
	Type            string
	TriggerTime     string
	DockerImageName string
	DockerImageTag  string
	ExecuteCommands []string
	SetupCommands   []string
	CachedDirs      []string
	Variables       []*ActionVariableEntity
	Shell           string
	CacheBaseImage  bool
	Services        []*ServiceEntity
}
type ActionVariableEntity struct {
	Key   string
	Value string
}
