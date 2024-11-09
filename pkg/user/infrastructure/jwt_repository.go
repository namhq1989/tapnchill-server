package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	appjwt "github.com/namhq1989/tapnchill-server/internal/jwt"
)

type JwtRepository struct {
	jwt appjwt.Operations
}

func NewJwtRepository(jwt appjwt.Operations) JwtRepository {
	return JwtRepository{
		jwt: jwt,
	}
}

func (r JwtRepository) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	return r.jwt.GenerateAccessToken(ctx, userID)
}
