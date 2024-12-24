package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type DowngradeExpiredSubscriptionsHandler struct {
	userRepository domain.UserRepository
}

func NewDowngradeExpiredSubscriptionsHandler(userRepository domain.UserRepository) DowngradeExpiredSubscriptionsHandler {
	return DowngradeExpiredSubscriptionsHandler{
		userRepository: userRepository,
	}
}

func (h DowngradeExpiredSubscriptionsHandler) DowngradeExpiredSubscriptions(ctx *appcontext.AppContext, _ domain.QueueDowngradeExpiredSubscriptionsPayload) error {
	ctx.Logger().Text("downgrade all expired subscriptions")
	_, err := h.userRepository.DowngradeAllExpiredSubscriptions(ctx)
	if err != nil {
		ctx.Logger().Error("failed to downgrade all expired subscriptions", err, appcontext.Fields{})
	}

	return err
}
