package main

import (
	"time"

	"github.com/prometheus/common/log"
	"github.com/sudarshan-reddy/airbender-appliance/handlers"
	"github.com/sudarshan-reddy/groove"

	"github.com/Sirupsen/logrus"
)

const (
	d4           = 4
	a0           = 14
	address      = 0x04
	timeInterval = 3 * time.Millisecond
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", err, msg)
	}
}

func main() {
	groveHandler, err := groove.InitGroove(address)
	defer groveHandler.Close()
	failOnError(err, "failed to initialise grove")
	handlers.TurnLEDOn(groveHandler, d4)
	time.Sleep(1 * time.Second)
	handlers.TurnLEDOff(groveHandler, d4)
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	for reading := range handlers.MonitorAirQuality(groveHandler, a0, ticker) {
		if reading.Err != nil {
			return
		}
		log.Infoln("sensor reading: ", reading.Reading)
	}
}
