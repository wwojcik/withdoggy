package gateway

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = 8080
)

func RunServer(ctx context.Context, logger *zap.Logger) error {
	// Listen an actual port.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer lis.Close()

	// Create a gRPC Server with gRPC interceptor.
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)
	logger.Sugar().Infof("server started on port: %d", port)
	return grpcServer.Serve(lis)
}
