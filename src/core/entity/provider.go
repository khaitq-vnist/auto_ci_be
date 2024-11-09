package entity

type ProviderEntity struct {
	BaseEntity
	Name    string
	Code    string
	ApiUrl  string
	AuthUrl string
}
