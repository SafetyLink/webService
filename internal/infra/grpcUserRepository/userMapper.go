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

func profileToModel(profile *authenticationv1.GetSelfResponse) *types.User {
	var chats []*types.Chat
	
	if len(profile.Chats) == 0 {
		chats = []*types.Chat{}
	} else {
		for _, c := range profile.Chats {
			chats = append(chats, &types.Chat{
				ChatID:         c.GetChatId(),
				UnreadMessages: c.GetUnreadMessage(),
				LastMessageAt:  c.GetLastMessageAt().AsTime(),
				Viewed:         c.GetViewed(),
				ViewedAt:       c.GetViewedAt().AsTime(),
				Users: &types.ChatUser{
					ID:       c.GetUserId(),
					Username: c.GetUsername(),
					AvatarID: c.GetAvatarId(),
				},
			})
		}
	}

	return &types.User{
		ID:        profile.GetUserId(),
		Username:  profile.GetUsername(),
		Email:     profile.GetEmail(),
		FirstName: profile.GetFirstname(),
		LastName:  profile.GetLastname(),
		AvatarID:  profile.GetAvatarId(),
		CreatedAt: profile.GetCreatedAt().AsTime(),
		UpdatedAt: profile.GetUpdatedAt().AsTime(),
		Chat:      chats,
	}
}
