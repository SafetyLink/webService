package grpcAuthenticationRepository

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"go.opentelemetry.io/otel/trace"
)

func (gs *GrpcAuthenticationRepo) Login(ctx context.Context, email, password string) (string, error) {
	ctx, span := gs.tracer.Start(ctx, "grpcAuthenticationRepository.Login", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	resp, err := gs.AuthGrpc.Login(ctx, &authenticationv1.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}
	return resp.JwtToken, nil
}

func (gs *GrpcAuthenticationRepo) Register(ctx context.Context, email, username, password string) (string, error) {
	ctx, span := gs.tracer.Start(ctx, "grpcAuthenticationRepository.Register", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

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
