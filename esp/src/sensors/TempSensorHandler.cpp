#include "TempSensorHandler.h"

#ifdef ESP32
    #include "dht22/dht22_sensor.h"
    DHT22Sensor sensor(2);
#elif ESP8266
    #include "sht31/sht31_sensor.h"
    SHT31Sensor sensor;
#endif

TempSensorHandler::TempSensorHandler() {}

void TempSensorHandler::initSensor() {
    sensor.begin();
}

float TempSensorHandler::readTemperature() {
    return sensor.readTemperature();
}

float TempSensorHandler::readHumidity() {
    return sensor.readHumidity();
}
