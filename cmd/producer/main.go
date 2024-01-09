package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
)

func main() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sc, err := stan.Connect("wb", "1234")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	log.Info("nats connection established")

	go func() {
		for {
			select {
			case sig := <-sigChan:
				log.Info("server stoped by signal", sig)
				os.Exit(1)
			}
		}
	}()

	dir, err := os.ReadDir("./testJSON")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {

		data, err := os.ReadFile(fmt.Sprintf("./testJSON/%s", file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		sc.Publish("orderWB", data)
		log.Info("the order has been uploaded")
		time.Sleep(time.Second * 4)
	}

}
