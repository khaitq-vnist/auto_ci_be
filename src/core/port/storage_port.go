package port

import "context"

type IStoragePort interface {
	UploadFile(ctx context.Context, fullPath, logData string) error
}
