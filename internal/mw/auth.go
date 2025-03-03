package mw

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Auth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata")
	}

	values := md.Get("x-auth")

	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "no metadata")
	}
	return handler(ctx, req)

}
