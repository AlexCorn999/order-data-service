package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan string)

	sc, err := stan.Connect("wb", "123")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sc.Subscribe("test", func(m *stan.Msg) {
		ch <- string(m.Data)
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case sig := <-sigChan:
			fmt.Println("stoped by signal", sig)
			return
		}
	}

}
