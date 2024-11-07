package shared

import "github.com/namhq1989/tapnchill-server/pkg/common/domain"

type Service struct {
	externalApiRepository domain.ExternalApiRepository
	cachingRepository     domain.CachingRepository
}

func NewService(externalApiRepository domain.ExternalApiRepository, cachingRepository domain.CachingRepository) Service {
	return Service{
		externalApiRepository: externalApiRepository,
		cachingRepository:     cachingRepository,
	}
}
