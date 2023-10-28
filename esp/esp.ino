#include "src/wifi/WiFiHandler.h"
#include "src/sensors/TempSensorHandler.h"

WiFiHandler wifiHandler;
TempSensorHandler tempSensorHandler;
int serialSpeed = 115200;

void setup() {
  Serial.begin(serialSpeed);

  wifiHandler.connectToWiFi();
  tempSensorHandler.initSensor();
}

void loop() {
    delay(5000);
	Serial.print(".");
}
