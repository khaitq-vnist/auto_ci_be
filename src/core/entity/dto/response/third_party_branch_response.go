package response

type ThirdPartyBranchResponse struct {
	Name   string `json:"name"`
	Commit Commit `json:"commit"`
}
type Commit struct {
	Sha string
	URL string
}
