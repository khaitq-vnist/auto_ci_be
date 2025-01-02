package properties

import "github.com/golibs-starter/golib/config"

type SonarProperties struct {
	BaseUrl     string
	AccessToken string
}

func (s SonarProperties) Prefix() string {
	return "app.services.sonar"
}

func NewSonarProperties(loader config.Loader) (*SonarProperties, error) {
	var properties SonarProperties
	err := loader.Bind(&properties)
	if err != nil {
		return nil, err
	}
	return &properties, nil
}
