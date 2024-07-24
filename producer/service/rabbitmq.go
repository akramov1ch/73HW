package service

import (
    "73HW/producer/model"

    "github.com/streadway/amqp"
)

type RabbitMQService interface {
    PublishMessage(msg model.Message) error
}

type rabbitMQService struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

func NewRabbitMQService() RabbitMQService {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        panic(err)
    }

    ch, err := conn.Channel()
    if err != nil {
        panic(err)
    }

    err = ch.ExchangeDeclare(
        "topic_logs",
        "topic",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        panic(err)
    }

    return &rabbitMQService{
        conn:    conn,
        channel: ch,
    }
}

func (s *rabbitMQService) PublishMessage(msg model.Message) error {
    return s.channel.Publish(
        "topic_logs",
        msg.RoutingKey,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(msg.Content),
        },
    )
}
