package grpc

import (
	"context"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"google.golang.org/grpc"
)

type server struct {
	hub Application
	userpb.UnimplementedUserServiceServer
}

var _ userpb.UserServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, hub *Application) error {
	userpb.RegisterUserServiceServer(registrar, server{hub: *hub})
	return nil
}

func (s server) GetHabitQuota(bgCtx context.Context, req *userpb.GetHabitQuotaRequest) (*userpb.GetHabitQuotaResponse, error) {
	return s.hub.GetHabitQuota(appcontext.NewGRPC(bgCtx), req)
}

func (s server) GetGoalQuota(bgCtx context.Context, req *userpb.GetGoalQuotaRequest) (*userpb.GetGoalQuotaResponse, error) {
	return s.hub.GetGoalQuota(appcontext.NewGRPC(bgCtx), req)
}

func (s server) GetTaskQuota(bgCtx context.Context, req *userpb.GetTaskQuotaRequest) (*userpb.GetTaskQuotaResponse, error) {
	return s.hub.GetTaskQuota(appcontext.NewGRPC(bgCtx), req)
}

func (s server) GetNoteQuota(bgCtx context.Context, req *userpb.GetNoteQuotaRequest) (*userpb.GetNoteQuotaResponse, error) {
	return s.hub.GetNoteQuota(appcontext.NewGRPC(bgCtx), req)
}
