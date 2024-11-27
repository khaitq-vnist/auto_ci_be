package response

type ProjectAnalyzeResponse struct {
	BuildTool Tool `json:"build_tool"`
}
type Tool struct {
	Maven  bool `json:"maven,omitempty"`
	Gradle bool `json:"gradle,omitempty"`
	NodeJs bool `json:"node.js,omitempty"`
	Go     bool `json:"go,omitempty"`
	Python bool `json:"python,omitempty"`
	Ruby   bool `json:"ruby,omitempty"`
}

func ToProjectAnalyzeResponse(tools map[string]bool) *ProjectAnalyzeResponse {
	return &ProjectAnalyzeResponse{
		BuildTool: Tool{
			Maven:  tools["Maven"],
			Gradle: tools["Gradle"],
			NodeJs: tools["Node.js"],
			Go:     tools["Go"],
			Python: tools["Python"],
			Ruby:   tools["Ruby"],
		},
	}
}
