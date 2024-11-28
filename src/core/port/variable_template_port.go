package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type IVariableTemplatePort interface {
	GetVariableTemplateByStageIDs(ctx context.Context, stageIDs []int64) ([]*entity.VariableTemplateEntity, error)
}
