package mqtt

import (
	"ESPeather/internal/db"
	"ESPeather/internal/models"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqqtMessage struct {
	Topic string
	Payload []byte
}

func connect(config *MqttConfig) mqtt.Client {
	opts := createClientOptions(config)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	token.Wait()

	if token.Error() != nil {
		fmt.Println("Connection error:", token.Error())
	}

	return client
}

func createClientOptions(config *MqttConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Address, config.Port))
	opts.SetClientID(config.ClientID)
	return opts
}

func Listen(config *MqttConfig) {
	client := connect(config)
	for _, topic := range config.Topics {
        client.Subscribe(topic, 0, handleMessage)
    }
}

func handleMessage(client mqtt.Client, msg mqtt.Message) {
	mqttMsg := parseMqttMessage(msg)
	topic := mqttMsg.Topic
	reading, err := unmarshalReading(mqttMsg.Payload)
	if err != nil {
		fmt.Println("Error unmarshalling MQTT payload:", err)
		return
	}
	db.InsertDB(topic, reading)
}

func parseMqttMessage(msg mqtt.Message) MqqtMessage {
	return MqqtMessage{
		Topic:   msg.Topic(),
		Payload: msg.Payload(),
	}
}

func unmarshalReading(payload []byte) (models.Reading, error) {
	var reading models.Reading
	err := json.Unmarshal(payload, &reading)
	return reading, err
}