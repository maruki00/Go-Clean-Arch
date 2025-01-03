package publisher

import (
	"context"

	"github.com/ClickHouse/ch-go"
	"github.com/rabbitmq/amqp091"
	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	AmqpConn  *amqp091.Connection
	QueueName string
}

func NewPublisher(con *amqp091.Connection, queuename string) *Publisher {
	return &Publisher{
		AmqpConn:  con,
		QueueName: queuename,
	}
}

func (p *Publisher) MakeChannel() (*amqp091.Channel, error) {
	ch, err := p.AmqpConn.Channel()
	if err != nil {
		panic(err)
	}
	queue, err := ch.Query(
		
	)
	return ch, nil
}

func (p *Publisher) Run(worker func(ctx context.Context, delivery <-chan amqp091.Delivery)) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	channel, err := p.MakeChannel()
	if err != nil {
		return err
	}
	channel.
}

func (p *Publisher) Publish(data []byte, contentType string) {

}
