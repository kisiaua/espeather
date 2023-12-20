#include "dht22_sensor.h"

DHT22Sensor::DHT22Sensor(int pin) : sensor(pin, DHT22) {}

void DHT22Sensor::begin() {
  sensor.begin();
}

float DHT22Sensor::readTemperature() {
  return sensor.readTemperature();
}

float DHT22Sensor::readHumidity() {
  return sensor.readHumidity();
}
