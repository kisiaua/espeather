#ifndef WIFI_HANDLER_H
#define WIFI_HANDLER_H

#ifdef ESP32
    #include <WiFi.h>
#elif ESP8266
    #include <ESP8266WiFi.h>
#endif

class WiFiHandler {
    public:
        WiFiHandler();
        void connectToWiFi();
        WiFiClient& getWiFiClient();
};

#endif
