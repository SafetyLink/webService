package message

import (
	"context"
	"github.com/SafetyLink/commons/types"
	"github.com/SafetyLink/webService/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
)

type Srv interface {
	CreateMessage(ctx context.Context, message types.Message) (*types.Message, error)
}

type Service struct {
	rabbitMQRepo repo.RabbitMQ
	tracer       trace.Tracer
}

func NewMessageSrv(rabbitMQ repo.RabbitMQ) Srv {
	return &Service{
		rabbitMQRepo: rabbitMQ,
	}
}
