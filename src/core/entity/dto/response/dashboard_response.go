package response

type DashboardResponse struct {
	TotalProjects     int64 `json:"total_projects"`
	TotalIntegrations int64 `json:"total_integrations"`
}
