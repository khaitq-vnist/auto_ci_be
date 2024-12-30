package request

type ThirdPartyCreateProjectRequest struct {
	DisplayName       string
	Name              string
	Integration       ThirdPartyCreateProjectIntegration
	ExternalProjectID string
}
type ThirdPartyCreateProjectIntegration struct {
	HashID string
}
