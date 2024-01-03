package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexCorn999/order-data-service/internal/apiserver"
	"github.com/AlexCorn999/order-data-service/internal/config"
	"github.com/AlexCorn999/order-data-service/internal/domain"
	"github.com/AlexCorn999/order-data-service/internal/nats"
	"github.com/nats-io/stan.go"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan []byte)

	// подключение к кластеру nats streaming
	sc, err := nats.NatsConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer (*sc).Close()

	// инициализация кэша
	// ?????

	// восстановление кэша из бд
	// ???????

	// подписка и чтение из канала
	(*sc).Subscribe("test", func(m *stan.Msg) {
		ch <- m.Data
	})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case data := <-ch:

				var order domain.Order
				if err := json.Unmarshal(data, &order); err != nil {
					log.Println("JSON - something wrong")
					continue
				}
				fmt.Printf("%+v", order)

			case sig := <-sigChan:
				fmt.Println("server stoped by signal", sig)
				os.Exit(1)
			}
		}
	}()

	// запуск сервера
	server := apiserver.NewAPIServer(config.NewConfig())
	if err := server.Start(ch); err != nil {
		log.Fatal(err)
	}

}
