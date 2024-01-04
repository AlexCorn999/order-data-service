package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexCorn999/order-data-service/internal/apiserver"
	"github.com/AlexCorn999/order-data-service/internal/config"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	server := apiserver.NewAPIServer(config.NewConfig())
	if err := server.Start(sigChan); err != nil {
		log.Fatal(err)
	}
}
