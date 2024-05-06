package main

import (
	"ESPeather/internal/api"
	"ESPeather/internal/mqtt"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	mqttConfig, err := mqtt.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		return
	}

	go mqtt.Listen(mqttConfig)

	go api.StartServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
