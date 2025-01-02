package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IServicePort interface {
	GetAllService(ctx context.Context) ([]*entity.ServiceEntity, error)
}
