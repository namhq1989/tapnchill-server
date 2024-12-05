package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
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

func (h GoogleSignInHandler) GoogleSignIn(ctx *appcontext.AppContext, performerID string, req dto.GoogleSignInRequest) (*dto.GoogleSignInResponse, error) {
	ctx.Logger().Info("new Google sign in request", appcontext.Fields{"performerID": performerID, "token": req.Token})

	ctx.Logger().Text("find current user")
	currentUser, err := h.userRepository.FindByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find current user", err, appcontext.Fields{})
		return nil, err
	}
	if currentUser == nil {
		ctx.Logger().ErrorText("current user not found, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("get user's data with Google token")
	ssoUser, err := h.ssoRepository.VerifyGoogleToken(ctx, req.Token)
	if err != nil {
		ctx.Logger().Error("failed to get staff data with Google token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("Google's user found, find user with provider id", appcontext.Fields{"id": ssoUser.UID})
	googleUser, err := h.userRepository.FindByAuthProviderID(ctx, ssoUser.UID)
	if err != nil {
		ctx.Logger().Error("failed to find user by email via grpc", err, appcontext.Fields{})
		return nil, err
	}
	if googleUser == nil {
		ctx.Logger().Text("user not found, add new provider to current user")
		currentUser.AddAuthProvider(domain.AuthProvider{
			Provider: domain.AuthProviderGoogle,
			ID:       ssoUser.UID,
			Name:     ssoUser.Name,
			Email:    ssoUser.Email,
		})

		ctx.Logger().Text("update user in db")
		if err = h.userRepository.Update(ctx, *currentUser); err != nil {
			ctx.Logger().Error("failed to persist user in db", err, appcontext.Fields{})
			return nil, err
		}
	} else {
		ctx.Logger().Text("user found, check current user & google user")
		if currentUser.ID != googleUser.ID {
			ctx.Logger().Text("google user is different from current user, set current user belongs to google user")
			if err = currentUser.SetBelongsTo(googleUser.ID); err != nil {
				ctx.Logger().Error("failed to set current user belongs to google user", err, appcontext.Fields{})
				return nil, err
			}
		}
	}

	ctx.Logger().Info("user found, generate access token", appcontext.Fields{"userID": currentUser.ID})
	accessToken, err := h.jwtRepository.GenerateAccessToken(ctx, currentUser.ID)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done Google sign in request")
	return &dto.GoogleSignInResponse{
		UserID:      currentUser.ID,
		Provider:    domain.AuthProviderGoogle,
		AccessToken: accessToken,
	}, nil
}
