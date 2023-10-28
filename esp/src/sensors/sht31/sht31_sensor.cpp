#include "sht31_sensor.h"

SHT31Sensor::SHT31Sensor() {}

void SHT31Sensor::begin() {
  sht31.begin(0x44);
}

float SHT31Sensor::readTemperature() {
  return sht31.readTemperature();
}

float SHT31Sensor::readHumidity() {
  return sht31.readHumidity();
}
