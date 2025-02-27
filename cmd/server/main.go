package main

import (
	"fmt"
	"net"
	"notes/internal/app/notes"
	"notes/internal/mw"
	desc "notes/pkg/api/notes/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			mw.Panic,
			mw.Logger,
			mw.Panic,
			mw.Validate,
			mw.Validate,
		),
	)

	reflection.Register(grpcServer)

	controller := notes.NewService()

	desc.RegisterNotesServer(grpcServer, controller)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
