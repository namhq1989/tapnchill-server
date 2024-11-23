package shared

import "github.com/namhq1989/tapnchill-server/pkg/user/domain"

type Service struct {
	userRepository    domain.UserRepository
	cachingRepository domain.CachingRepository
}

func NewService(userRepository domain.UserRepository, cachingRepository domain.CachingRepository) Service {
	return Service{
		userRepository:    userRepository,
		cachingRepository: cachingRepository,
	}
}
