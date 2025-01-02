package properties

import "github.com/golibs-starter/golib/config"

type BuddyProperties struct {
	BaseUrl     string
	AccessToken string
	Workspace   string
}

func (b BuddyProperties) Prefix() string {
	return "app.services.buddy"
}

func NewBuddyProperties(loader config.Loader) (*BuddyProperties, error) {
	props := &BuddyProperties{}
	err := loader.Bind(props)
	return props, err
}
