package adapters

import (
	"encoding/json"
	"hex_sub/src/payments/domain"

	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	
}


func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {

	conn, err := amqp.Dial("amqp://carlos:carlos@100.27.181.40/")
	if err != nil {
		log.Printf("Error conectando a RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error abriendo canal: %v", err)
		return nil, err
	}
	
	_, err = ch.QueueDeclare(
		"pagos",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error declarando la cola: %v", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, data domain.Payment) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error convirtiendo evento a JSON: %v", err)
		return err
	}

	err = r.ch.Publish(
		"",
		"pagos",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Error enviando mensaje a RabbitMQ: %v", err)
		return err
	}

	log.Println("Evento publicado:", eventType)
	return nil
}