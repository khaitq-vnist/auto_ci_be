package response

type SonarCreateTokenResponse struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Type  string `json:"type"`
}
