package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	data, err := os.ReadFile("./model.json")
	if err != nil {
		log.Fatal(err)
	}

	sc, err := stan.Connect("wb", "1234")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	log.Println("connected to nats")

	go func() {
		for {
			select {
			case sig := <-sigChan:
				fmt.Println("server stoped by signal", sig)
				os.Exit(1)
			}
		}
	}()

	// добавить отправку множества значений в канал
	for i := 0; i < 5; i++ {
		sc.Publish("orderWB", data)
		time.Sleep(time.Second * 4)
	}
}
