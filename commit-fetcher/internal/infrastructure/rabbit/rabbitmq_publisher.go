package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewPublisher(rabbitMQURL, repoName string) (*Publisher, error) {
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

	return &Publisher{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (p *Publisher) PublishSignal(message string) error {
	err := p.channel.Publish(
		"",           // exchange
		p.queue.Name, // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	log.Printf("Signal sent: %s", message)
	return nil
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
