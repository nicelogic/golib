package producer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/nicelogic/pubsub/event"
)

type Producer struct {
	Producer pulsar.Producer
}

func (producer *Producer) SendAsync(ctx context.Context,
	userIds []string,
	payload any) error {

	event := &event.Event{
		UserIds: userIds,
		Payload: payload,
	}
	payloadJson, err := json.Marshal(event)
	if err != nil {
		log.Printf("json marshal err(%v), but still return true(ntf not affect whether success)\n", err)
		return err
	}
	log.Printf("SendAsync payload(%v) success\n", payload)
	producer.Producer.SendAsync(ctx,
		&pulsar.ProducerMessage{Payload: payloadJson},
		func(mi pulsar.MessageID, pm *pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Printf("pulsar send ntf(%v), err(%v)\n", payload, err)
			}
			log.Printf("pulsar send ntf(%v), success, msgId(%v)\n", payload, mi)
		})
	return nil
}
