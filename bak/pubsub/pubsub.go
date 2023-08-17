package pubsub

import (
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/nicelogic/config"
	"github.com/nicelogic/pubsub/producer"
)

type Pubsub struct {
	Client            pulsar.Client
	DefaultEventTopic string
}

func (pubsub *Pubsub) Init(configFilePath string) (err error) {
	type ClientConfig struct {
		Url                 string
		Operation_timeout   int
		Connection_timeout  int
		Default_event_topic string
	}
	clientConfig := ClientConfig{}
	err = config.Init(configFilePath, &clientConfig)
	if err != nil {
		log.Printf("config init err(%v)\n", err)
		return err
	}
	log.Printf("(%#v)\n", clientConfig)

	pubsub.Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:               clientConfig.Url,
		OperationTimeout:  time.Duration(clientConfig.Operation_timeout) * time.Second,
		ConnectionTimeout: time.Duration(clientConfig.Connection_timeout) * time.Second,
	})
	if err != nil {
		log.Printf("Could not instantiate Pulsar client(%v)", err)
		return err
	}
	if clientConfig.Default_event_topic == "" {
		err = fmt.Errorf("clientConfig.Default_event_topic is empty, not config")
		return err
	}
	pubsub.DefaultEventTopic = clientConfig.Default_event_topic
	log.Printf("pubsub init success\n")
	return err
}

func (pubsub *Pubsub) CreateProducer(option pulsar.ProducerOptions) (*producer.Producer, error) {
	pulsarProducer, err := pubsub.Client.CreateProducer(option)
	if err != nil {
		log.Printf("pubsub.Client.CreateProducer err(%v)\n", err)
		return nil, err
	}
	producer := &producer.Producer{
		Producer: pulsarProducer,
	}
	return producer, nil
}
