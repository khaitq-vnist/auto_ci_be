package properties

import "github.com/golibs-starter/golib/config"

type GithubProperties struct {
	BaseUrl string
}

func (g GithubProperties) Prefix() string {
	return "app.services.github"
}

func NewGithubProperties(loader config.Loader) (*GithubProperties, error) {
	props := &GithubProperties{}
	err := loader.Bind(props)
	return props, err
}
