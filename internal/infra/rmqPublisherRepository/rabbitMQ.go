package rmqPublisherRepository

import (
	"context"
	"encoding/json"
	"github.com/SafetyLink/commons/types"
	"github.com/SafetyLink/webService/internal/domain/repo"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQPubRepository struct {
	rabbitMQChannel *amqp.Channel
	rabbitMQQueue   amqp.Queue
	logger          *zap.Logger
}

func NewRabbitMQPubRepository(logger *zap.Logger, rabbitMQConn *amqp.Connection) repo.RabbitMQ {
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		logger.Error("failed to create rabbitMQ channel", zap.Error(err))
		return nil
	}

	rmqQueue, err := ch.QueueDeclare(
		"messages", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		logger.Error("failed to declare a rabbitMQ queue ", zap.Error(err))
		return nil
	}

	return &RabbitMQPubRepository{
		rabbitMQChannel: ch,
		rabbitMQQueue:   rmqQueue,
		logger:          logger,
	}
}

func (rr *RabbitMQPubRepository) PublishMessage(ctx context.Context, message types.Message) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		rr.logger.Error("failed to marshal message to JSON", zap.Error(err))
		return err
	}

	err = rr.rabbitMQChannel.PublishWithContext(ctx,
		"",
		rr.rabbitMQQueue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         jsonMessage,
		})
	if err != nil {
		rr.logger.Error("failed to publish message", zap.Error(err))
		return err
	}

	return nil

}
