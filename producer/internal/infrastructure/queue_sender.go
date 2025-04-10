package infrastructure

import (
	"encoding/json"
	"go-api-ddd/producer/internal/domain"
	"net"
	"os"
)


func SendToQueue(msg domain.Message) error {
	conn, err := net.Dial("tcp", os.Getenv("MQ_HOST")+":5672")

	if err != nil {
		return err
	}

	defer conn.Close()

	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	_, err = conn.Write(data)

	return err
}