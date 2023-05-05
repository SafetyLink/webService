package clients

import (
	"fmt"
	"github.com/SafetyLink/webService/internal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GrpcAuthenticationClient(logger *zap.Logger, config *internal.Config) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.Services.AuthenticationService.Port, grpc.WithInsecure())
	if err != nil {
		logger.Panic(fmt.Sprintf("failed to connect to %s", config.Services.AuthenticationService.Name))
	}
	logger.Info(fmt.Sprintf("connected to %s", config.Services.AuthenticationService.Name))

	return conn
}
