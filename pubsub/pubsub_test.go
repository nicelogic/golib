package pubsub_test

import (
	"github.com/nicelogic/pubsub"
	"testing"
)

func TestPubsub(t *testing.T) {
	pubsub := &pubsub.Pubsub{}
	err := pubsub.Init("./config-pulsar.yml")
	if err != nil {
		t.Errorf("pubsubinit error")
	}
}
