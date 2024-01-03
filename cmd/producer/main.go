package main

import (
	"log"
	"os"

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

	log.Println("connected")

	sc.Publish("test", data)

}
