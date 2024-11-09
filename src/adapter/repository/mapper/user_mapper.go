package mapper

import (
	model2 "github.com/khaitq-vnist/auto_ci_be/adapter/repository/postgres/model"
	"github.com/khaitq-vnist/auto_ci_be/core/entity"
)

func ToUserModel(userEntity *entity.UserEntity) *model2.UserModel {
	if userEntity == nil {
		return nil
	}
	return &model2.UserModel{
		BaseModel: model2.BaseModel{
			ID: userEntity.ID,
		},
		Email:    userEntity.Email,
		Name:     userEntity.Name,
		Password: userEntity.Password,
	}
}

func ToUserEntity(userModel *model2.UserModel) *entity.UserEntity {
	if userModel == nil {
		return nil
	}
	return &entity.UserEntity{
		BaseEntity: entity.BaseEntity{
			ID:        userModel.ID,
			CreatedAt: userModel.CreatedAt.Unix(),
			UpdatedAt: userModel.UpdatedAt.Unix(),
		},
		Email:    userModel.Email,
		Name:     userModel.Name,
		Password: userModel.Password,
	}
}
