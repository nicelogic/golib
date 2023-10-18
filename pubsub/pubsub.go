package pubsub

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nicelogic/pubsub/event"
)

type Pubsub struct {
	Client            *nats.Conn
	DefaultEventTopic string
}

func (pubsub *Pubsub) Init(url string, defaultTopic string) (err error) {
	type ClientConfig struct {
		Url                 string
		Default_event_topic string
	}
	clientConfig := ClientConfig{
		Url: url,
		Default_event_topic: defaultTopic,
	}
	log.Printf("(%#v)\n", clientConfig)
	if clientConfig.Default_event_topic == "" {
		err = fmt.Errorf("clientConfig.Default_event_topic is empty, not config")
		return err
	}
	pubsub.DefaultEventTopic = clientConfig.Default_event_topic

	nc, err := nats.Connect(clientConfig.Url,
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("DisconnectErrHandler nats.Conn(%v), err(%v)\n", nc, err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("ReconnectHandler nats.Conn(%v)\n", nc)
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("client closed, nats.Conn(%v)\n", nc)
		}),
		nats.ErrorHandler(func(nc *nats.Conn, subscription *nats.Subscription, err error) {
			if subscription != nil {
				log.Printf("ErrorHandler, Async error in %q/%q: %v", subscription.Subject, subscription.Queue, err)
			} else {
				log.Printf("ErrorHandler, Async error outside subscription: %v", err)
			}
		}),
	)
	if err != nil {
		log.Printf("connect error(%v)", err)
		return err
	}
	pubsub.Client = nc
	// defer nc.Drain()

	log.Printf("pubsub init success\n")
	return err
}

func (pubsub *Pubsub) SendAsync(
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
	if err := pubsub.Client.Publish(pubsub.DefaultEventTopic, payloadJson); err != nil {
		log.Printf("publish error(%v)\n", err)
		return err
	}
	log.Printf("SendAsync payload(%v) success\n", payload)
	return nil
}
