package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initRPC() *grpc.Server {
	s := grpc.NewServer(
		grpc.Creds(nil),
	)
	reflection.Register(s)
	return s
}
