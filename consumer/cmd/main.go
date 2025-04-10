package main

import (
	"go-api-ddd/consumer/internal/infrastructure"
	"go-api-ddd/consumer/internal/infrastructure/repositories"
)

func main() {
    repositories.InitDB()
    infrastructure.StartListener()
}
