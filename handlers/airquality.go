package handlers

import (
	"fmt"
	"time"

	"github.com/sudarshan-reddy/groove"
)

// AirSensorResponse is the response given by MonitorAirQuality
type AirSensorResponse struct {
	Reading int
	Err     error
}

// MonitorAirQuality takes in the handler, pin and
// monitor interval and returns it to a channel
func MonitorAirQuality(grove groove.Handler, pin byte,
	ticker *time.Ticker) chan AirSensorResponse {
	airResponse := make(chan AirSensorResponse)
	fmt.Println("air quality monitor starting")

	go func() {
		defer close(airResponse)
		for range ticker.C {
			reading, err := grove.AnalogRead(pin)
			airResponse <- AirSensorResponse{
				Reading: reading,
				Err:     err,
			}
		}
	}()
	return airResponse
}
