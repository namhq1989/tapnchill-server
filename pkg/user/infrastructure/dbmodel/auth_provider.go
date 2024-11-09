package dbmodel

import "github.com/namhq1989/tapnchill-server/pkg/user/domain"

type AuthProvider struct {
	Provider string `bson:"provider"`
	ID       string `bson:"id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
}

func (ap AuthProvider) ToDomain() domain.AuthProvider {
	return domain.AuthProvider{
		Provider: ap.Provider,
		ID:       ap.ID,
		Name:     ap.Name,
		Email:    ap.Email,
	}
}

func (AuthProvider) FromDomain(provider domain.AuthProvider) AuthProvider {
	return AuthProvider{
		Provider: provider.Provider,
		ID:       provider.ID,
		Name:     provider.Name,
		Email:    provider.Email,
	}
}
