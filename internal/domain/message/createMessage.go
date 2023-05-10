package message

import (
	"context"
	"github.com/SafetyLink/commons/snowflake"
	"github.com/SafetyLink/commons/types"
)

func (ms *Service) CreateMessage(ctx context.Context, message types.Message) (*types.Message, error) {
	ctx, span := ms.tracer.Start(ctx, "messageService.createMessage")
	defer span.End()

	message.MessageID = snowflake.GenerateSnowflakeID()

	err := ms.rabbitMQRepo.PublishMessage(ctx, message)
	if err != nil {
		return &types.Message{}, err
	}
	return &message, nil
}
