package main

import (
	"fmt"
	"go-api-ddd/producer/internal/interfaces"
	"net/http"
)

func main() {
	http.HandleFunc("/send", interfaces.SendMessageHandler)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}