// Temporary solution to faulty sensor that gives 'nan' readings after some time

#include "RestartHandler.h"
#include <Arduino.h>

bool RestartHandler::checkAndRestart(float temperature, float humidity) {
    if (isnan(temperature) || isnan(humidity)) {
        Serial.println("Invalid reading detected. Restarting...");
        ESP.restart();
        return true;
    }
    return false;
}