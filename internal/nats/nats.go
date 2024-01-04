package nats

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

type NatsST struct {
	Sc           stan.Conn
	subscription stan.Subscription
	subject      string
	durableName  string
	Data         chan []byte
}

// NatsConnection устанавливает соединение с кластером в nats streaming.
func NewNatsST(clusterID, clientID string) (*NatsST, error) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return nil, err
	}

	return &NatsST{
		Sc:          sc,
		subject:     "orderWB",
		durableName: "my-durable",
		Data:        make(chan []byte),
	}, nil
}

func (n *NatsST) SubscribeCh() error {
	handler := func(m *stan.Msg) {
		n.Data <- m.Data
	}
	sub, err := n.Sc.Subscribe(n.subject, handler, stan.DurableName(n.durableName))
	if err != nil {
		return fmt.Errorf("can't subscribe: %w", err)
	}
	n.subscription = sub
	return nil
}

// UnsubsribeNs отписка от получения данных.
func (n *NatsST) UnsubsribeNs() error {
	return n.subscription.Unsubscribe()
}

// Close закрывает подключение к nats.
func (n *NatsST) Close() error {
	return n.Sc.Close()
}
