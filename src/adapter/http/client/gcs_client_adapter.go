package client

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/properties"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"google.golang.org/api/option"
	"time"
)

type GCSClientAdapter struct {
	props *properties.GCSProperties
}

func (g GCSClientAdapter) GenerateURLDownload(ctx context.Context, fullPath string) (string, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(g.props.CredentialJSON)))
	if err != nil {
		log.Error(ctx, "Error while creating GCloud client", err)
		return "", err
	}
	defer func(client *storage.Client) {
		err := client.Close()
		if err != nil {
			log.Error(ctx, "Error while closing GCloud client", err)
		}
	}(client)

	// Parse CredentialJSON to extract the GoogleAccessID and PrivateKey
	var credentials struct {
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
	}
	if err := json.Unmarshal([]byte(g.props.CredentialJSON), &credentials); err != nil {
		log.Error(ctx, "Error parsing credentials JSON", err)
		return "", err
	}

	// Define the signed URL options
	signedURLOptions := &storage.SignedURLOptions{
		GoogleAccessID: credentials.ClientEmail,
		PrivateKey:     []byte(credentials.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(15 * time.Minute), // Link valid for 15 minutes
	}

	// Generate the signed URL
	signedURL, err := storage.SignedURL(g.props.Bucket, fullPath, signedURLOptions)
	if err != nil {
		log.Error(ctx, "Error while generating signed URL", err)
		return "", err
	}

	return signedURL, nil
}

func (g GCSClientAdapter) UploadFile(ctx context.Context, fullPath, logData string) (string, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(g.props.CredentialJSON)))
	if err != nil {
		log.Error(ctx, "Error while creating GCloud client", err)
		return "", err
	}
	defer func(client *storage.Client) {
		err := client.Close()
		if err != nil {
			log.Error(ctx, "Error while closing GCloud client", err)
		}
	}(client)

	// Upload the file
	bucket := client.Bucket(g.props.Bucket)
	object := bucket.Object(g.props.ProjectID + "/" + fullPath)
	writer := object.NewWriter(ctx)
	writer.ContentType = "text/plain"
	if _, err = writer.Write([]byte(logData)); err != nil {
		log.Error(ctx, "Error while writing file to GCloud", err)
		return "", err
	}
	if err = writer.Close(); err != nil {
		log.Error(ctx, "Error while closing writer", err)
		return "", err
	}
	log.Info(ctx, "Uploaded file to GCloud successfully")

	// Parse CredentialJSON to extract GoogleAccessID and PrivateKey
	var credentials struct {
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
	}
	if err := json.Unmarshal([]byte(g.props.CredentialJSON), &credentials); err != nil {
		log.Error(ctx, "Error parsing credentials JSON", err)
		return "", err
	}

	// Generate the signed URL
	signedURLOptions := &storage.SignedURLOptions{
		GoogleAccessID: credentials.ClientEmail,
		PrivateKey:     []byte(credentials.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(7 * 24 * time.Hour), // Link valid for 15 minutes
	}

	signedURL, err := storage.SignedURL(g.props.Bucket, g.props.ProjectID+"/"+fullPath, signedURLOptions)
	if err != nil {
		log.Error(ctx, "Error while generating signed URL", err)
		return "", err
	}

	return signedURL, nil
}

func NewGCSClientAdapter(props *properties.GCSProperties) port.IStoragePort {
	return &GCSClientAdapter{props: props}
}
