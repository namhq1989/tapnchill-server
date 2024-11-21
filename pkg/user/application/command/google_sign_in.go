package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type GoogleSignInHandler struct {
	userRepository domain.UserRepository
	ssoRepository  domain.SSORepository
	jwtRepository  domain.JwtRepository
}

func NewGoogleSignInHandler(userRepository domain.UserRepository, ssoRepository domain.SSORepository, jwtRepository domain.JwtRepository) GoogleSignInHandler {
	return GoogleSignInHandler{
		userRepository: userRepository,
		jwtRepository:  jwtRepository,
		ssoRepository:  ssoRepository,
	}
}

func (h GoogleSignInHandler) GoogleSignIn(ctx *appcontext.AppContext, req dto.GoogleSignInRequest) (*dto.GoogleSignInResponse, error) {
	ctx.Logger().Info("new Google sign in request", appcontext.Fields{"token": req.Token})

	ctx.Logger().Text("get user's data with Google token")
	googleUser, err := h.ssoRepository.VerifyGoogleToken(ctx, req.Token)
	if err != nil {
		ctx.Logger().Error("failed to get staff data with Google token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("Google's user found, find application's user with email via grpc", appcontext.Fields{"email": googleUser.Email})
	user, err := h.userRepository.FindByEmail(ctx, googleUser.Email)
	if err != nil {
		ctx.Logger().Error("failed to find user by email via grpc", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, create new one")
		user, err = domain.NewGoogleUser(googleUser.UID, googleUser.Name, googleUser.Email)
		if err != nil {
			ctx.Logger().Error("failed to create new user", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist user in db")
		if err = h.userRepository.Create(ctx, *user); err != nil {
			ctx.Logger().Error("failed to persist user in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Info("user found, generate access token", appcontext.Fields{"userID": user.ID})
	accessToken, err := h.jwtRepository.GenerateAccessToken(ctx, user.ID)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done Google sign in request")
	return &dto.GoogleSignInResponse{
		AccessToken: accessToken,
	}, nil
}
