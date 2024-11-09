package postgres

import "github.com/khaitq-vnist/auto_ci_be/core/port"

type UserRepositoryAdapter struct {
	*BaseRepository
}

func NewUserRepositoryAdapter(base *BaseRepository) port.IUserPort {
	return &UserRepositoryAdapter{
		base,
	}
}
