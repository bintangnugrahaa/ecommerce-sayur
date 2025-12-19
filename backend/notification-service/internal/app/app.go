package app

import (
	"notification-service/config"
	"notification-service/internal/adapter/message"
	"notification-service/internal/adapter/rabbitmq"
	"notification-service/utils"

	"github.com/labstack/echo/v4"
)

func RunServer() {
	cfg := config.NewConfig()
	emailMessage := message.NewMessageEmail(cfg)
	rabbitMQAdapter := rabbitmq.NewConsumeRabbitMQ(emailMessage)

	e := echo.New()

	go func() {
		err := rabbitMQAdapter.ConsumeMessage(utils.NOTIF_EMAIL_VERIFICATION)
		if err != nil {
			e.Logger.Errorf("Failed to consume RabbitMQ for %s: %v", utils.NOTIF_EMAIL_VERIFICATION, err)
		}
	}()

	go func() {
		err := rabbitMQAdapter.ConsumeMessage(utils.NOTIF_EMAIL_FORGOT_PASSWORD)
		if err != nil {
			e.Logger.Errorf("Failed to consume RabbitMQ for %s: %v", utils.NOTIF_EMAIL_FORGOT_PASSWORD, err)
		}
	}()

	e.Logger.Fatal(e.Start(":" + cfg.App.AppPort))
}
