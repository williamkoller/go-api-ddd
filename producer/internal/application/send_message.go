package application

import (
	"go-api-ddd/producer/internal/domain"
	"go-api-ddd/producer/internal/infrastructure"
)


func SendMessage(msg domain.Message) error {
	return infrastructure.SendToQueue(msg)
}
