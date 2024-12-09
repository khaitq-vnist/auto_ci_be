package properties

import "github.com/golibs-starter/golib/config"

type GCSProperties struct {
	ProjectID      string
	Bucket         string
	CredentialJSON string
}

func (G GCSProperties) Prefix() string {
	return "app.services.gcs"
}

func NewGCSProperties(loader config.Loader) (*GCSProperties, error) {
	props := &GCSProperties{}
	err := loader.Bind(props)
	return props, err
}
