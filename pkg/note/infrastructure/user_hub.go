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

func (r UserHub) GetNoteQuota(ctx *appcontext.AppContext, userID string) (int64, bool, error) {
	resp, err := r.client.GetNoteQuota(ctx.Context(), &userpb.GetNoteQuotaRequest{
		TraceId: ctx.GetTraceID(),
		UserId:  userID,
	})
	if err != nil {
		return 0, true, err
	}

	return resp.GetLimit(), resp.GetIsFree(), nil
}
