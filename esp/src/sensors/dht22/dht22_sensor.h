#ifndef DHT22_SENSOR_H
#define DHT22_SENSOR_H

#include <DHT.h>

class DHT22Sensor {
public:
  DHT22Sensor(int pin);
  void begin();
  float readTemperature();
  float readHumidity();

private:
  DHT sensor;
};

#endif
