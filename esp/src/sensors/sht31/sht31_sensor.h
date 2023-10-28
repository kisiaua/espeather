#ifndef SHT31_SENSOR_H
#define SHT31_SENSOR_H

#include <Adafruit_SHT31.h>

class SHT31Sensor {
public:
  SHT31Sensor();
  void begin();
  float readTemperature();
  float readHumidity();

private:
  Adafruit_SHT31 sht31;
};

#endif
