package main

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
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

	// добавить отправку множества значений в канал
	for i := 0; i < 5; i++ {
		sc.Publish("orderWB", data)
		time.Sleep(time.Second * 4)
	}
}
