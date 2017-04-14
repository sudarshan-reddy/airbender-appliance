package main

import (
	"time"

	"github.com/sudarshan-reddy/airbender-appliance/handlers"
	"github.com/sudarshan-reddy/groove"

	"github.com/Sirupsen/logrus"
)

const (
	d4      = 4
	a1      = 1
	address = 0x04
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
}
