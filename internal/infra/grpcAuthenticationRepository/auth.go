package grpcAuthenticationRepository

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type GrpcAuthenticationRepo struct {
	AuthGrpc authenticationv1grpc.AuthenticationServiceClient
	tracer   trace.Tracer
}

func NewGrpcAuthenticationRepository(authGrpcConn *grpc.ClientConn, tracer trace.Tracer) repo.Authentication {
	return &GrpcAuthenticationRepo{
		AuthGrpc: authenticationv1grpc.NewAuthenticationServiceClient(authGrpcConn),
		tracer:   tracer,
	}
}
