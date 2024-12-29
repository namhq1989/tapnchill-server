package grpcclient

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newConn(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
}

func NewUserClient(_ *appcontext.AppContext, addr string) (userpb.UserServiceClient, error) {
	conn, err := newConn(addr)
	if err != nil {
		return nil, err
	}

	return userpb.NewUserServiceClient(conn), nil
}
