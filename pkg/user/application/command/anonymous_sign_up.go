package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type AnonymousSignUpHandler struct {
	userRepository domain.UserRepository
	jwtRepository  domain.JwtRepository
}

func NewAnonymousSignUpHandler(userRepository domain.UserRepository, jwtRepository domain.JwtRepository) AnonymousSignUpHandler {
	return AnonymousSignUpHandler{
		userRepository: userRepository,
		jwtRepository:  jwtRepository,
	}
}

func (h AnonymousSignUpHandler) AnonymousSignUp(ctx *appcontext.AppContext, req dto.AnonymousSignUpRequest) (*dto.AnonymousSignUpResponse, error) {
	ctx.Logger().Info("new anonymous sign up request", appcontext.Fields{"clientID": req.ClientID, "checksum": req.Checksum})

	if !h.userRepository.ValidateAnonymousChecksum(ctx, req.ClientID, req.Checksum) {
		ctx.Logger().Text("invalid checksum, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("checksum is valid, create new user model")
	user, err := domain.NewUser(req.ClientID)
	if err != nil {
		ctx.Logger().Error("failed to create new user model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist user in db")
	if err = h.userRepository.Create(ctx, *user); err != nil {
		ctx.Logger().Error("failed to persist user in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("generate access token")
	token, err := h.jwtRepository.GenerateAccessToken(ctx, user.ID)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("return anonymous sign up response")
	return &dto.AnonymousSignUpResponse{
		AccessToken: token,
	}, nil
}
