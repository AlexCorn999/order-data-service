package nats

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

type NatsST struct {
	sc           stan.Conn
	subscription stan.Subscription
	subject      string
	durableName  string
	Data         chan []byte
}

// NatsConnection устанавливает соединение с кластером в nats streaming.
// Returns an object with connection data, subject, durable name and initializes the channel to which the data will be sent.
func NewNatsST(clusterID, clientID string) (*NatsST, error) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	return &NatsST{
		sc:          sc,
		subject:     "orderWB",
		durableName: "my-durable",
		Data:        make(chan []byte),
	}, nil
}

// Subscribe makes a subscription with the given parameters to a cluster in nats.
func (n *NatsST) SubscribeCh() error {
	handler := func(m *stan.Msg) {
		n.Data <- m.Data
	}
	sub, err := n.sc.Subscribe(n.subject, handler, stan.DurableName(n.durableName))
	if err != nil {
		return fmt.Errorf("can't subscribe: %w", err)
	}
	n.subscription = sub
	return nil
}

// UnsubsribeNs removes interest in subscribing to the subject.
func (n *NatsST) UnsubsribeNs() error {
	return n.subscription.Unsubscribe()
}

// Close closes the connection to the cluster in nats.
func (n *NatsST) Close() error {
	return n.sc.Close()
}
