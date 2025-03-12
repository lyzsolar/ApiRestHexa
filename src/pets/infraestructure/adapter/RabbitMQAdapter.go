package adapter

import (
	"APIDOS/src/pets/application/repository"
	"APIDOS/src/pets/domain/entities"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repository.IMessageService = (*RabbitMQAdapter)(nil)

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("")
	if err != nil {
		log.Println("Error connecting to RabbitMQ:", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error opening channel:", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"pets", // Queue name
		true,   // Durable
		false,  // Auto-delete
		false,  // Exclusive
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Println("Error declaring queue:", err)
		return nil, err
	}

	err = ch.Confirm(false)
	if err != nil {
		log.Println("Error enabling message confirmations:", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, pets entities.Mascotas) error {
	body, err := json.Marshal(pets)
	if err != nil {
		log.Println("Error converting event to JSON:", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",     // Exchange
		"pets", // Routing key (queue name)
		true,   // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error sending message to RabbitMQ:", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Message confirmed by RabbitMQ")
	case <-nack:
		log.Println("Message was not confirmed")
	}

	log.Println("Event published:", eventType)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Error closing RabbitMQ channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error closing RabbitMQ connection: %v", err)
	}
}
