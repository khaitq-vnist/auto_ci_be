package postgres

import "github.com/khaitq-vnist/auto_ci_be/core/port"

type IntegrationRepositoryAdapter struct {
	*BaseRepository
}

func NewIntegrationRepositoryAdapter(base *BaseRepository) port.IIntegrationPort {
	return &IntegrationRepositoryAdapter{
		base,
	}
}
