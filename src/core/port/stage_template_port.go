package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IStageTemplatePort interface {
	GetStageTemplateByIds(ctx context.Context, IDs []int64) ([]*entity.StageTemplateEntity, error)
	GetStageTemplateByPipelineTemplateID(ctx context.Context, pipelineTemplateID int64) ([]*entity.StageTemplateEntity, error)
}
