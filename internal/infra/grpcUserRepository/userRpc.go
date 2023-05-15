package grpcUserRepository

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"github.com/SafetyLink/commons/types"
)

func (gs *GrpcUserRepo) GetUserByID(ctx context.Context, profileID int64) (*types.User, error) {
	userResp, err := gs.UserGrpc.GetUserByID(ctx, &authenticationv1.GetUserByIDRequest{
		UserId: profileID,
	})
	if err != nil {
		return nil, err
	}

	return userIDToModel(userResp), nil
}

func (gs *GrpcUserRepo) GetSelf(ctx context.Context) (*types.User, error) {
	profileResp, err := gs.UserGrpc.GetSelf(ctx, &authenticationv1.GetSelfRequest{UserId: ctx.Value("userID").(int64)})
	if err != nil {
		return nil, err
	}

	return profileToModel(profileResp), nil
}
