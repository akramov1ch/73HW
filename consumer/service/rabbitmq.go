package service

import (
    "log"

    "github.com/streadway/amqp"
)

type RabbitMQService interface {
    ConsumeMessages(bindingKey string, queueName string)
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

func (s *rabbitMQService) ConsumeMessages(bindingKey string, queueName string) {
    q, err := s.channel.QueueDeclare(
        queueName,
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        panic(err)
    }

    err = s.channel.QueueBind(
        q.Name,
        bindingKey,
        "topic_logs",
        false,
        nil,
    )
    if err != nil {
        panic(err)
    }

    msgs, err := s.channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        panic(err)
    }

    for msg := range msgs {
        log.Printf("Received a message: %s", msg.Body)
    }
}
