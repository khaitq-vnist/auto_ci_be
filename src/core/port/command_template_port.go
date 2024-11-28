package port

import (
	"context"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

type ICommandTemplatePort interface {
	GetCommandTemplateByStageIDs(ctx context.Context, ids []int64) ([]*entity.CommandTemplateEntity, error)
}
