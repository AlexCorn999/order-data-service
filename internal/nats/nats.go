package nats

import (
	"github.com/nats-io/stan.go"
)

// NatsConnection устанавливает соединение с кластером в nats streaming.
func NatsConnection() (*stan.Conn, error) {
	sc, err := stan.Connect("wb", "1")
	if err != nil {
		return nil, err
	}
	return &sc, nil
}
