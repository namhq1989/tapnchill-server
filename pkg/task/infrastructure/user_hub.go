package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
)

type UserHub struct {
	client userpb.UserServiceClient
}

func NewUserHub(client userpb.UserServiceClient) UserHub {
	return UserHub{
		client: client,
	}
}

func (r UserHub) GetGoalQuota(ctx *appcontext.AppContext, userID string) (int64, error) {
	resp, err := r.client.GetGoalQuota(ctx.Context(), &userpb.GetGoalQuotaRequest{
		TraceId: ctx.GetTraceID(),
		UserId:  userID,
	})
	if err != nil {
		return 0, err
	}

	return resp.GetLimit(), nil
}

func (r UserHub) GetTaskQuota(ctx *appcontext.AppContext, userID string) (int64, error) {
	resp, err := r.client.GetTaskQuota(ctx.Context(), &userpb.GetTaskQuotaRequest{
		TraceId: ctx.GetTraceID(),
		UserId:  userID,
	})
	if err != nil {
		return 0, err
	}

	return resp.GetLimit(), nil
}
