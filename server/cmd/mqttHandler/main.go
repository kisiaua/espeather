package main

import (
	"ESPeather/internal/mqtt"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	mqttConfig, err := mqtt.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
        return
	}

	go mqtt.Listen(mqttConfig)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}