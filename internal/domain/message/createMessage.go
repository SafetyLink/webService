package message

import (
	"context"
	"github.com/SafetyLink/commons/snowflake"
	"github.com/SafetyLink/commons/types"
)

func (ms *Service) CreateMessage(ctx context.Context, message types.Message) (*types.Message, error) {
	message.MessageID = snowflake.GenerateSnowflakeID()

	err := ms.rabbitMQRepo.PublishMessage(ctx, message)
	if err != nil {
		return &types.Message{}, err
	}
	return &message, nil
}
