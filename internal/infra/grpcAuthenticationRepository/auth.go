package grpcAuthenticationRepository

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"

	"context"
	"google.golang.org/grpc"
)

type GrpcAuthenticationRepo struct {
	AuthRepo authenticationv1grpc.AuthenticationServiceClient
}

func NewGrpcAuthenticationRepository(conn *grpc.ClientConn) *GrpcAuthenticationRepo {
	authRepo := authenticationv1grpc.NewAuthenticationServiceClient(conn)

	return &GrpcAuthenticationRepo{
		AuthRepo: authRepo,
	}
}

func (gs GrpcAuthenticationRepo) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := gs.AuthRepo.Login(ctx, &authenticationv1.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}
	return resp.JwtToken, nil

}

func (gs GrpcAuthenticationRepo) Register(ctx context.Context, email, username, password string) (string, error) {
	resp, err := gs.AuthRepo.Register(ctx, &authenticationv1.RegisterRequest{
		Email:    email,
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}
	return resp.JwtToken, nil

}
