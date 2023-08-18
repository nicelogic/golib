package pubsub_test

import (
	"log"
	"sync"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/nicelogic/pubsub"
)

func TestPubsub(t *testing.T) {
	pubsub := &pubsub.Pubsub{}
	err := pubsub.Init("./config-nats.yml")
	if err != nil {
		t.Errorf("pubsubinit error")
	}

	pubsub.SendAsync([]string{"user0"}, "test")
}

func TestPubsub_subscribe(t *testing.T) {
	pubsub := &pubsub.Pubsub{}
	err := pubsub.Init("./config-nats.yml")
	if err != nil {
		t.Errorf("pubsubinit error")
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	sub, err := pubsub.Client.Subscribe("event", func(m *nats.Msg) {
		log.Printf("Received a message(%s)\n", string(m.Data))
		// wg.Done()
	})
	if err != nil {
		log.Printf("subscribe err(%v)\n", err)
		return
	}
	defer sub.Drain()

	// Wait for a message to come in
	wg.Wait()
}
