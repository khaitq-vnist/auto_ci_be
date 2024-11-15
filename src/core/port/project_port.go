package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IProjectPort interface {
	SaveProject(ctx context.Context, entity *entity.ProjectEntity) (*entity.ProjectEntity, error)
	GetListProjectByUserId(ctx context.Context, userId int64) ([]*entity.ProjectEntity, error)
}
