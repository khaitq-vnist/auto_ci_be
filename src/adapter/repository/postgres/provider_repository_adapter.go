package postgres

import (
	"context"
	"errors"
	"github.com/golibs-starter/golib/log"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/mapper"
	"github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/common"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
	"gorm.io/gorm"
)

type ProviderRepositoryAdapter struct {
	*BaseRepository
}

func (p ProviderRepositoryAdapter) GetProviderByCode(ctx context.Context, code string) (*entity.ProviderEntity, error) {
	var providerModel model.ProviderModel
	if err := p.db.WithContext(ctx).Model(&model.ProviderModel{}).Where("code = ?", code).First(&providerModel).Error; err != nil {
		log.Error(ctx, "Error when get provider by code", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(common.ErrRecordNotFound)
		}
		return nil, err
	}
	return mapper.ToProviderEntity(&providerModel), nil
}

func NewProviderRepositoryAdapter(base *BaseRepository) port.IProviderPort {
	return &ProviderRepositoryAdapter{
		base,
	}
}
