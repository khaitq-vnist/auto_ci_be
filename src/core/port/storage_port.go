package port

import "context"

type IStoragePort interface {
	UploadFile(ctx context.Context, fullPath, logData string) (string, error)
	GenerateURLDownload(ctx context.Context, fullPath string) (string, error)
}
