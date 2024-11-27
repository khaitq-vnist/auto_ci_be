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

type GithubContentResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
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
func ToThirdPartyContentResponse(content *GithubContentResponse) *response.ThirdPartyContentResponse {
	return &response.ThirdPartyContentResponse{
		Name: content.Name,
		Type: content.Type,
		Path: content.Path,
	}
}
func ToListThirdPartyContentResponse(content []*GithubContentResponse) []*response.ThirdPartyContentResponse {
	var result []*response.ThirdPartyContentResponse
	for _, c := range content {
		result = append(result, ToThirdPartyContentResponse(c))
	}
	return result
}

type GitHubBranchResponse struct {
	Name   string `json:"name"`
	Commit struct {
		SHA string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	Protected  bool `json:"protected"`
	Protection struct {
		RequiredStatusChecks struct {
			EnforcementLevel string   `json:"enforcement_level"`
			Contexts         []string `json:"contexts"`
		} `json:"required_status_checks"`
	} `json:"protection"`
	ProtectionURL string `json:"protection_url"`
}

func ToThirdPartyBranchResponse(branch *GitHubBranchResponse) *response.ThirdPartyBranchResponse {
	return &response.ThirdPartyBranchResponse{
		Name: branch.Name,
		Commit: response.Commit{
			Sha: branch.Commit.SHA,
			URL: branch.Commit.URL,
		},
	}
}

func ToListThirdPartyBranchResponse(branches []*GitHubBranchResponse) []*response.ThirdPartyBranchResponse {
	var result []*response.ThirdPartyBranchResponse
	for _, b := range branches {
		result = append(result, ToThirdPartyBranchResponse(b))
	}
	return result
}
