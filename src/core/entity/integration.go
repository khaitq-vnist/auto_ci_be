package entity

type IntegrationEntity struct {
	BaseEntity
	Name             string
	UserId           int64
	ProviderId       int64
	ProviderName     string
	ProviderCode     string
	AccessToken      string
	ProviderUsername string
	ThirdPartyHashId string
}
