package grpcUserRepository

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"github.com/SafetyLink/commons/types"
)

func userIDToModel(user *authenticationv1.GetUserByIDResponse) *types.User {
	return &types.User{
		ID:        user.GetUserId(),
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		FirstName: user.GetFirstname(),
		LastName:  user.GetLastname(),
		AvatarID:  user.GetAvatarId(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}

}
