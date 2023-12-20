#ifndef MQTT_HANDLER_H
#define MQTT_HANDLER_H

#ifdef ESP32
    #include <WiFi.h>
#elif ESP8266
    #include <ESP8266WiFi.h>
#endif

#include <PubSubClient.h>

class mqttHandler {
    public:
        mqttHandler();
        void initMQTT(WiFiClient&);
        void setServer();
        bool connected();
        void reconnect();
        void publish(const char*, String);

    private:
        PubSubClient client;
};

#endif
