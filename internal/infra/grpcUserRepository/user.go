package grpcUserRepository

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"google.golang.org/grpc"
)

type GrpcUserRepo struct {
	UserGrpc authenticationv1grpc.UserServiceClient
}

func NewGrpcUserRepository(userGrpc *grpc.ClientConn) repo.User {
	return &GrpcUserRepo{
		UserGrpc: authenticationv1grpc.NewUserServiceClient(userGrpc),
	}
}
