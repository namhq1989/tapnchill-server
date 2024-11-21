package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/sso"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type SSORepository struct {
	sso sso.Operations
}

func NewSSORepository(sso sso.Operations) SSORepository {
	return SSORepository{
		sso: sso,
	}
}

func (r SSORepository) VerifyGoogleToken(ctx *appcontext.AppContext, token string) (*domain.SSOGoogleUser, error) {
	user, err := r.sso.VerifyGoogleToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &domain.SSOGoogleUser{
		UID:   user.UID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
