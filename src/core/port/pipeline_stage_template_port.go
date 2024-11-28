package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IPipelineStageTemplatePort interface {
	GetPipelineStageTemplateByTemplateID(ctx context.Context, templateID int64) ([]*entity.PipelineStageTemplateEntity, error)
}
