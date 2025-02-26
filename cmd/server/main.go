package main

import (
	"fmt"
	"net"
	"notes/internal/app/notes"
	desc "notes/pkg/api/notes/v1"

	"google.golang.org/grpc"
)

const (
	grpcPort = 50051
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	controller := notes.NewService()

	desc.RegisterNotesServer(grpcServer, controller)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
