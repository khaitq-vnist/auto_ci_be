package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IIntegrationPort interface {
	SaveIntegration(ctx *context.Context, integration *entity.IntegrationEntity) (*entity.IntegrationEntity, error)
	GetIntegrationByUserId(ctx *context.Context, userId int64) ([]*entity.IntegrationEntity, error)
}
