#ifndef TEMP_SENSOR_HANDLER_H
#define TEMP_SENSOR_HANDLER_H

class TempSensorHandler {
    public:
        TempSensorHandler();
        void initSensor();
        float readTemperature();
        float readHumidity();
};

#endif
