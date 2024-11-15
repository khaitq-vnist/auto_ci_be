package response

import (
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"time"
)

type GithubUserInfoResponse struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type GithubRepoInfo struct {
	ID          int64     `json:"id"`
	NodeID      string    `json:"node_id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Private     bool      `json:"private"`
	HtmlUrl     string    `json:"html_url"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *GithubUserInfoResponse) ToUserInfo() *response.ThirdPartyProviderUserInfoResponse {
	return &response.ThirdPartyProviderUserInfoResponse{
		Username: u.Login,
		Name:     u.Name,
		Company:  u.Company,
		Email:    u.Email,
	}
}
func ToThirdPartyProviderRepoResponse(repo *GithubRepoInfo) *response.ThirdPartyProviderReposResponse {
	return &response.ThirdPartyProviderReposResponse{
		ID:          repo.ID,
		Name:        repo.Name,
		Private:     repo.Private,
		FullName:    repo.FullName,
		Description: repo.Description,
		HtmlUrl:     repo.HtmlUrl,
		CreateAt:    repo.CreateAt.Unix(),
		UpdatedAt:   repo.UpdatedAt.Unix(),
	}
}
func ToListThirdPartyProviderRepoResponse(repos []*GithubRepoInfo) []*response.ThirdPartyProviderReposResponse {
	var result []*response.ThirdPartyProviderReposResponse
	for _, repo := range repos {
		result = append(result, &response.ThirdPartyProviderReposResponse{
			ID:        repo.ID,
			Name:      repo.Name,
			Private:   repo.Private,
			FullName:  repo.FullName,
			CreateAt:  repo.CreateAt.Unix(),
			UpdatedAt: repo.UpdatedAt.Unix(),
		})
	}
	return result
}
