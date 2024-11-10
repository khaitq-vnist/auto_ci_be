package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IProviderPort interface {
	GetProviderByCode(ctx context.Context, code string) (*entity.ProviderEntity, error)
}
