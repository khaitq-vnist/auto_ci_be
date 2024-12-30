package entity

type ServiceEntity struct {
	BaseEntity
	Type       string
	Version    string
	Connection ServiceConnectionEntity
}
type ServiceConnectionEntity struct {
	Host     string
	DB       string
	Port     string
	User     string
	Password string
}
