package entity

type ProjectEntity struct {
	BaseEntity
	Name                string
	FullName            string
	Private             bool
	OwnerID             uint
	Description         *string
	Language            *string
	ProviderRepoID      int64
	ProviderCode        string
	ProviderUsername    string
	HtmlUrl             string
	IntegrationId       int64
	ThirdPartyProjectID string
}
