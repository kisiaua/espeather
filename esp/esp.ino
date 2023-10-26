#include "src/wifi/WiFiHandler.h"

WiFiHandler wifiHandler;
int serialSpeed = 115200;

void setup() {
  Serial.begin(serialSpeed);

  wifiHandler.connectToWiFi();
}

// the loop function runs over and over again forever
void loop() {
    delay(5000);
	Serial.print(".");
}
