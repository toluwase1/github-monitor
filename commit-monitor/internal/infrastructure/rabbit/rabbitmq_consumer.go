package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewConsumer(rabbitMQURL, repoName string) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queueName := fmt.Sprintf("commit_monitor_queue_%s", repoName)
	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (c *Consumer) ConsumeSignals() (<-chan amqp.Delivery, error) {
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (c *Consumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
