package repo

import (
	"context"
	"github.com/SafetyLink/commons/types"
)

type Authentication interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, username, password string) (string, error)
}

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*types.User, error)
	GetProfile(ctx context.Context, userID int64) (*types.User, error)
}
