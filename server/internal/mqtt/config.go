package mqtt

import (
	"encoding/json"
	"os"
)

type MqttConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	ClientID string `json:"clientID"`
	Topics   []string `json:"topics"`
}

func LoadConfig() (*MqttConfig, error) {
	configStr := os.Getenv("MQTT_CONFIG")
	
	var mqttConfig MqttConfig
	err := json.Unmarshal([]byte(configStr), &mqttConfig)
	if err != nil {
		return nil, err
	}

	return &mqttConfig, nil
}