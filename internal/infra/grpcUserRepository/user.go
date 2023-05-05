package grpcUserRepository

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"fmt"
	"github.com/SafetyLink/commons/types"
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

func (gs GrpcUserRepo) GetUserByID(ctx context.Context, userID int64) (*types.User, error) {
	userResp, err := gs.UserGrpc.GetUserByID(ctx, &authenticationv1.GetUserByIDRequest{
		UserId: userID,
	})

	if err != nil {
		fmt.Println(err)
		return &types.User{}, err
	}

	return userIDToModel(userResp), nil

}
