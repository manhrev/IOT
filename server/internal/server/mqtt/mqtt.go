package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/manhrev/IOT/server/internal/repository"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Payload struct {
	Value int
}

type MqttServer struct {
	mqtt.Client
	repository *repository.Repository
}

const (
	USERNAME = "admin"
	PASSWORD = "admin"
)

func New(host, port, clientID string, entClient *ent.Client) *MqttServer {
	broker := fmt.Sprintf("tcp://%s:%s", host, port)
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("server_client").SetUsername(USERNAME).SetPassword(PASSWORD)

	repo := repository.New(entClient)

	opts.SetKeepAlive(60 * time.Second)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		// repo.Data.Create(context.Background(), msg.Topic(), int64(msg.Payload()[]))
		var value Payload
		err := json.Unmarshal(msg.Payload(), &value)
		if err != nil {
			log.Printf("Data is currently with the wrong type")
		}

		repo.Data.Create(context.Background(), msg.Topic(), int64(value.Value), timestamppb.New(time.Now()))

		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	// opts.SetPingTimeout(5 * time.Second)
	return &MqttServer{
		mqtt.NewClient(opts),
		repo,
	}
}

func (m *MqttServer) Start() error {
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
