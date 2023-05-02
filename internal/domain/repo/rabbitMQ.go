package repo

import (
	"context"
	"github.com/SafetyLink/commons/types"
)

type RabbitMQ interface {
	PublishMessage(ctx context.Context, message types.Message) error
}
