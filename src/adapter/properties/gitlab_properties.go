package properties

import "github.com/golibs-starter/golib/config"

type GitlabProperties struct {
	BaseUrl string
}

func (g GitlabProperties) Prefix() string {
	return "app.services.gitlab"
}

func NewGitlabProperties(loader config.Loader) (*GitlabProperties, error) {
	props := &GitlabProperties{}
	err := loader.Bind(props)
	return props, err
}
