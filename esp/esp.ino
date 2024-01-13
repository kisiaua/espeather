#include "src/wifi/WiFiHandler.h"
#include "src/sensors/TempSensorHandler.h"
#include "src/mqtt/mqttHandler.h"
#include "src/utils/RestartHandler.h"

WiFiHandler wifiHandler;
TempSensorHandler tempSensorHandler;
mqttHandler mqttHandler;
RestartHandler restartHandler;
int serialSpeed = 115200;

void setup() {
    Serial.begin(serialSpeed);

    wifiHandler.connectToWiFi();
    tempSensorHandler.initSensor();
    mqttHandler.initMQTT(wifiHandler.getWiFiClient());
    mqttHandler.setServer();
}

void loop() {

    if (!mqttHandler.connected()) {
        mqttHandler.reconnect();
    }

    float temperature = tempSensorHandler.readTemperature();
    float humidity = tempSensorHandler.readHumidity();
    if (restartHandler.checkAndRestart(temperature, humidity)) {
      return;
    }
    
    String jsonReadings = "{\"temperature\":" + String(temperature, 2) + ",\"humidity\":" + String(humidity, 2) + "}";

    mqttHandler.publish(jsonReadings);

    delay(30000);
}
