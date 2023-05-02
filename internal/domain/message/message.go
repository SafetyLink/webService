package message

import (
	"context"
	"github.com/SafetyLink/commons/types"
	"github.com/SafetyLink/webService/internal/domain/repo"
)

type Srv interface {
	CreateMessage(ctx context.Context, message types.Message) (*types.Message, error)
}

type Service struct {
	rabbitMQRepo repo.RabbitMQ
}

func NewMessageSrv(rabbitMQ repo.RabbitMQ) Srv {
	return &Service{
		rabbitMQRepo: rabbitMQ,
	}
}
