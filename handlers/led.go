package handlers

import (
	"time"

	"github.com/sudarshan-reddy/groove"
)

// TurnLEDOn turns the LED on for the grovePi's mentioned
// digital pin
func TurnLEDOn(handler groove.Handler, pin byte) error {
	err := handler.PinMode(pin, "output")
	if err != nil {
		return err
	}
	handler.DigitalWrite(pin, 1)
	time.Sleep(100 * time.Millisecond)
	return nil
}

// TurnLEDOff turns the LED off for the grovePi's mentioned
// digital pin
func TurnLEDOff(handler groove.Handler, pin byte) error {
	err := handler.PinMode(pin, "output")
	if err != nil {
		return err
	}
	handler.DigitalWrite(pin, 0)
	time.Sleep(100 * time.Millisecond)
	return nil
}
