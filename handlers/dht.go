package handlers

import (
	"fmt"
	"time"

	"github.com/sudarshan-reddy/groove"
)

// DHTResponse is the response struct for MonitorDHT
type DHTResponse struct {
	Temperature float32
	Humidity    float32
	Err         error
}

// MonitorDHTQuality takes in the handler, pin and
// monitor interval and returns it to a channel
func MonitorDHTQuality(done chan struct{}, grove groove.Handler, pin byte,
	ticker *time.Ticker) chan DHTResponse {
	dhtResponse := make(chan DHTResponse)
	fmt.Println("air quality monitor starting")

	go func() {
		defer close(dhtResponse)
		for range ticker.C {
			select {
			case <-done:
				return
			default:
				temp, humidity, err := grove.ReadDHT(pin)
				dhtResponse <- DHTResponse{
					Temperature: temp,
					Humidity:    humidity,
					Err:         err,
				}
			}
		}
	}()
	return dhtResponse
}
