package client

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"google.golang.org/api/option"
)

func NewGCloudClient(props *properties.GCSProperties) *storage.Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(props.CredentialJSON)))
	if err != nil {
		log.Panic("Error while creating GCloud client", err)
	}
	return client
}
