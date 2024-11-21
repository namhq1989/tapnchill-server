package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type AnonymousSignInHandler struct {
	userRepository  domain.UserRepository
	jwtRepository   domain.JwtRepository
	queueRepository domain.QueueRepository
}

func NewAnonymousSignInHandler(
	userRepository domain.UserRepository,
	jwtRepository domain.JwtRepository,
	queueRepository domain.QueueRepository,
) AnonymousSignInHandler {
	return AnonymousSignInHandler{
		userRepository:  userRepository,
		jwtRepository:   jwtRepository,
		queueRepository: queueRepository,
	}
}

func (h AnonymousSignInHandler) AnonymousSignIn(ctx *appcontext.AppContext, req dto.AnonymousSignInRequest) (*dto.AnonymousSignInResponse, error) {
	ctx.Logger().Info("new anonymous sign in request", appcontext.Fields{"clientID": req.ClientID, "source": req.Source, "checksum": req.Checksum})

	if !h.userRepository.ValidateAnonymousChecksum(ctx, req.ClientID, req.Checksum) {
		ctx.Logger().Text("invalid checksum, respond")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("checksum is valid, create new user model")
	user, err := domain.NewUser(req.ClientID, req.Source)
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

	ctx.Logger().Text("enqueue tasks")
	if err = h.enqueueTasks(ctx, *user); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Text("return anonymous sign in response")
	return &dto.AnonymousSignInResponse{
		AccessToken: token,
	}, nil
}

func (h AnonymousSignInHandler) enqueueTasks(ctx *appcontext.AppContext, user domain.User) error {
	ctx.Logger().Text("add task create user default Goal")
	if err := h.queueRepository.CreateUserDefaultGoal(ctx, domain.QueueCreateUserDefaultGoalPayload{
		UserID: user.ID,
	}); err != nil {
		return err
	}

	return nil
}
