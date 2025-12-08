package message

import (
	"encoding/json"
	"fmt"
	"order-service/config"
	"order-service/internal/core/domain/entity"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

type PublishRabbitMQInterface interface {
	PublishUpdateStock(productID int64, quantity int64)
	PublishOrderToQueue(order entity.OrderEntity) error
	PublishSendEmailUpdateStatus(email, message string) error
	PublishUpdateStatus(queuename string, orderID int64, status string) error
}

type PublishRabbitMQ struct {
	cfg *config.Config
}

// PublishUpdateStatus implements PublishRabbitMQInterface.
func (p *PublishRabbitMQ) PublishUpdateStatus(queuename string, orderID int64, status string) error {
	conn, err := p.cfg.NewRabbitMQ()
	if err != nil {
		log.Errorf("[PublishUpdateStatus-1] Failed to connect to RabbitMQ: %v", err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[PublishUpdateStatus-2] Failed to open a channel: %v", err)
		return err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queuename,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("[PublishUpdateStatus-3] Failed to declare a queue: %v", err)
		return err
	}

	orderStatus := map[string]string{
		"orderID": fmt.Sprintf("%d", orderID),
		"status":  status,
	}

	body, err := json.Marshal(orderStatus)
	if err != nil {
		log.Errorf("[PublishUpdateStatus-4] Failed to marshal JSON: %v", err)
		return err
	}

	return ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// PublishSendEmailUpdateStatus implements PublishRabbitMQInterface.
func (p *PublishRabbitMQ) PublishSendEmailUpdateStatus(email string, message string) error {
	conn, err := p.cfg.NewRabbitMQ()
	if err != nil {
		log.Errorf("[PublishSendEmailUpdateStatus-1] Failed to connect to RabbitMQ: %v", err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[PublishSendEmailUpdateStatus-2] Failed to open a channel: %v", err)
		return err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		p.cfg.PublisherName.EmailUpdateStatus,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("[PublishSendEmailUpdateStatus-3] Failed to declare a queue: %v", err)
		return err
	}

	notification := map[string]string{
		"email":   email,
		"message": message,
	}

	body, err := json.Marshal(notification)
	if err != nil {
		log.Errorf("[PublishSendEmailUpdateStatus-4] Failed to marshal JSON: %v", err)
		return err
	}

	return ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// PublishOrderToQueue implements PublishRabbitMQInterface.
func (p *PublishRabbitMQ) PublishOrderToQueue(order entity.OrderEntity) error {
	conn, err := p.cfg.NewRabbitMQ()
	if err != nil {
		log.Errorf("[PublishOrderToQueue-1] Failed to connect to RabbitMQ: %v", err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[PublishOrderToQueue-2] Failed to open a channel: %v", err)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		p.cfg.PublisherName.OrderPublish,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("[PublishOrderToQueue-3] Failed to declare queue: %v", err)
		return err
	}

	data, _ := json.Marshal(order)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		log.Errorf("[PublishOrderToQueue-4] Failed to publish message: %v", err)
		return err
	}

	return nil
}

// PublishUpdateStock implements PublishRabbitMQInterface.
func (p *PublishRabbitMQ) PublishUpdateStock(productID int64, quantity int64) {
	conn, err := p.cfg.NewRabbitMQ()
	if err != nil {
		log.Errorf("[PublishUpdateStock-1] Failed to connect to RabbitMQ: %v", err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[PublishUpdateStock-2] Failed to open a channel: %v", err)
		return
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(p.cfg.PublisherName.ProductUpdateStock, true, false, false, false, nil)
	if err != nil {
		log.Errorf("[PublishUpdateStock-3] Failed to declare a queue: %v", err)
		return
	}

	order := entity.PublishOrderItemEntity{
		ProductID: productID,
		Quantity:  quantity,
	}

	data, err := json.Marshal(order)
	if err != nil {
		log.Errorf("[PublishUpdateStock-4] Failed to marshal order: %v", err)
		return
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		log.Errorf("[PublishUpdateStock-5] Failed to publish message: %v", err)
		return
	}

	log.Info("Pesan order dikirim ke RabbitMQ")
}

func NewPublisherRabbitMQ(cfg *config.Config) PublishRabbitMQInterface {
	return &PublishRabbitMQ{cfg: cfg}
}
