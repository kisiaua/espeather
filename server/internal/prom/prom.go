package prom

import (
	"ESPeather/internal/models"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    temperature = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "temperature_readings",
            Help: "Current temperature reading.",
        },
        []string{"location"},
    )
    humidity = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "humidity_readings",
            Help: "Current relative humidity reading.",
        },
        []string{"location"},
    )
)

func init() {
    prometheus.MustRegister(temperature)
    prometheus.MustRegister(humidity)
}

func RecordMetrics(reading models.Reading) {
    temperature.WithLabelValues("living_room").Set(float64(reading.Temperature))
    humidity.WithLabelValues("living_room").Set(float64(reading.Humidity))
}

func Expose() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}