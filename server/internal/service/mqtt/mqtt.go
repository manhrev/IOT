package mqtt

import (
	"context"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/manhrev/IOT/server/internal/repository"
	"github.com/manhrev/IOT/server/pkg/ent"
)

type MqttService struct {
	mqtt.Client
	repository *repository.Repository
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func New(host, port, clientID string, entClient *ent.Client) *MqttService {
	broker := fmt.Sprintf("tcp://%s:%s", host, port)
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("server_client")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(5 * time.Second)
	return &MqttService{
		mqtt.NewClient(opts),
		repository.New(entClient),
	}
}

func (m *MqttService) Start() error {
	if token := m.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	// Subscribe to topics
	feeds, err := m.repository.Feed.List(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		if token := m.Subscribe(feed.FeedName, 0, nil); token.Wait() && token.Error() != nil {
			return token.Error()
		}
		log.Printf("Subscribe to topic: %s", feed.FeedName)
	}

	return nil
}
