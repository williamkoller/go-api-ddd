package application

import (
	"go-api-ddd/consumer/internal/domain"
	"go-api-ddd/consumer/internal/infrastructure/repositories"
)

func ProcessMessage(msg domain.Message) error {
    return repositories.SaveMessageToDB(msg)
}
