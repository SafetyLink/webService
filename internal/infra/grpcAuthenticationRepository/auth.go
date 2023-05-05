package grpcAuthenticationRepository

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"google.golang.org/grpc"

	"context"
)

type GrpcAuthenticationRepo struct {
	AuthGrpc authenticationv1grpc.AuthenticationServiceClient
}

func NewGrpcAuthenticationRepository(authGrpcConn *grpc.ClientConn) repo.Authentication {
	return &GrpcAuthenticationRepo{
		AuthGrpc: authenticationv1grpc.NewAuthenticationServiceClient(authGrpcConn),
	}
}

func (gs GrpcAuthenticationRepo) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := gs.AuthGrpc.Login(ctx, &authenticationv1.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}
	return resp.JwtToken, nil

}

func (gs GrpcAuthenticationRepo) Register(ctx context.Context, email, username, password string) (string, error) {
	resp, err := gs.AuthGrpc.Register(ctx, &authenticationv1.RegisterRequest{
		Email:    email,
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}
	return resp.JwtToken, nil

}
