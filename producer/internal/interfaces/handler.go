package interfaces

import (
	"encoding/json"
	"fmt"
	"go-api-ddd/producer/internal/application"
	"go-api-ddd/producer/internal/domain"
	"net/http"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request)  {
	var msg domain.Message

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := application.SendMessage(msg); err != nil {
		fmt.Println("Erro ao enviar para fila:", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}