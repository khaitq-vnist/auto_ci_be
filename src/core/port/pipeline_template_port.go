package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IPipelineTemplatePort interface {
	GetPipelineTemplateByBuildTools(ctx context.Context, typePipeline string) (*entity.PipelineTemplateEntity, error)
}
