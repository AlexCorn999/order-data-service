package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("wb", "123456789")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	log.Println("connected")

	for i := 0; ; i++ {
		sc.Publish("test", []byte("Hello"+strconv.Itoa(i)))
		time.Sleep(time.Second * 2)
	}
}
