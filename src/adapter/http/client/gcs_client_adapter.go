package client

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"google.golang.org/api/option"
)

type GCSClientAdapter struct {
	props *properties.GCSProperties
}

func (g GCSClientAdapter) UploadFile(ctx context.Context, fullPath, logData string) error {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(g.props.CredentialJSON)))
	if err != nil {
		log.Error(ctx, "Error while creating GCloud client", err)
		return err
	}
	defer func(client *storage.Client) {
		err := client.Close()
		if err != nil {
			log.Error(ctx, "Error while closing GCloud client", err)
		}
	}(client)
	bucket := client.Bucket(g.props.Bucket)
	object := bucket.Object(g.props.ProjectID + "/" + fullPath)
	writer := object.NewWriter(ctx)
	writer.ContentType = "text/plain"
	if _, err = writer.Write([]byte(logData)); err != nil {
		log.Error(ctx, "Error while writing file to GCloud", err)
		return err
	}
	if err = writer.Close(); err != nil {
		log.Error(ctx, "Error while closing writer", err)
		return err
	}
	log.Info(ctx, "Upload file to GCloud successfully")
	return nil
}

func NewGCSClientAdapter(props *properties.GCSProperties) port.IStoragePort {
	return &GCSClientAdapter{props: props}
}
