package shared

import "github.com/namhq1989/tapnchill-server/pkg/user/domain"

type Service struct {
	userRepository        domain.UserRepository
	cachingRepository     domain.CachingRepository
	externalAPIRepository domain.ExternalAPIRepository
}

func NewService(
	userRepository domain.UserRepository,
	cachingRepository domain.CachingRepository,
	externalAPIRepository domain.ExternalAPIRepository,
) Service {
	return Service{
		userRepository:        userRepository,
		cachingRepository:     cachingRepository,
		externalAPIRepository: externalAPIRepository,
	}
}
