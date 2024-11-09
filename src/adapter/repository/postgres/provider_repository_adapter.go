package postgres

import "github.com/khaitq-vnist/auto_ci_be/core/port"

type ProviderRepositoryAdapter struct {
	*BaseRepository
}

func NewProviderRepositoryAdapter(base *BaseRepository) port.IProviderPort {
	return &ProviderRepositoryAdapter{
		base,
	}
}
