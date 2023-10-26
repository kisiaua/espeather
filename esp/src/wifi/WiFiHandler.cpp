#include "WiFiHandler.h"
#include "WiFiCredentials.h"

#ifdef ESP32
    #include <WiFi.h>
#elif ESP8266
    #include <ESP8266WiFi.h>
#endif

WiFiHandler::WiFiHandler() {}

void WiFiHandler::connectToWiFi() {
    WiFi.begin(ssid, password);

    Serial.println("Connecting to WiFi");
    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.println(".");
    }
    Serial.println("Connected to WiFi");
}
