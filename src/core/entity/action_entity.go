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
type ServiceEntity struct {
	Type       string
	Version    string
	Connection ServiceConnectionEntity
}
type ServiceConnectionEntity struct {
	Host     string
	DB       string
	Port     int
	User     string
	Password string
}
