package response

type ThirdPartyProviderUserInfoResponse struct {
	Username string
	Name     string
	Company  string
	Email    string
}

type ThirdPartyProviderReposResponse struct {
	ID          int64
	Name        string
	Private     bool
	FullName    string
	Description string
	Language    string
	HtmlUrl     string
	CreateAt    int64
	UpdatedAt   int64
}

type ThirdPartyContentResponse struct {
	Name string
	Type string
	Path string
}
