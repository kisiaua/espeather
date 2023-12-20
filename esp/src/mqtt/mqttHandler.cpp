#include "mqttHandler.h"
#include "mqttConfig.h"

#ifdef ESP32
    #include <WiFi.h>
#elif ESP8266
    #include <ESP8266WiFi.h>
#endif

#include <PubSubClient.h>

mqttHandler::mqttHandler() {}

void mqttHandler::initMQTT(WiFiClient& wifiClient) {
    client.setClient(wifiClient);
}

void mqttHandler::setServer() {
    client.setServer(mqtt_server, mqtt_port);
}

bool mqttHandler::connected() {
    return client.connected();
}

void mqttHandler::reconnect() {
    while (!connected()) {
        Serial.print("Connecting to MQTT broker...");
        if (client.connect(client_id)) {
          Serial.println("Connected to MQTT broker");
        } else {
          Serial.println("Couldn't connect to MQTT broker. Retrying in 5 seconds...");
          delay(5000);
        }
        delay(3000);
    }
}

void mqttHandler::publish(const char* topic, String message) {
    client.publish(topic, message.c_str());
}