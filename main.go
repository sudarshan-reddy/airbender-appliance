package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sudarshan-reddy/airbender-appliance/handlers"
	"github.com/sudarshan-reddy/airbender-appliance/mq"
	"github.com/sudarshan-reddy/groove"
	"github.com/sudarshan-reddy/groove/dht"

	log "github.com/Sirupsen/logrus"
)

const (
	d3           = 3
	d4           = 4
	a0           = 14
	address      = 0x04
	timeInterval = 119 * time.Second
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func main() {
	cfg, err := loadConfigs()
	failOnError(err, "failed to load configs")
	zone, err := time.LoadLocation(cfg.TimeZone)
	failOnError(err, "error loading location")
	f, err := os.OpenFile(cfg.LogFileName, os.O_WRONLY|os.O_CREATE, 0755)
	failOnError(err, "failed to create log file")
	defer f.Close()
	log.SetOutput(f)

	groveHandler, err := groove.InitGroove(address)
	failOnError(err, "failed to initialise grove")
	defer groveHandler.Close()

	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	done := make(chan struct{})

	mqttClient, err := mq.NewClient(cfg.MQTTTopic, cfg.MQTTURL, cfg.MQTTClient)
	failOnError(err, "failed to load client")
	defer mqttClient.Close()

	for reading := range handlers.MonitorAirQuality(done, groveHandler, a0, ticker) {
		message, err := prepMessage(zone, reading.Reading, cfg.ApplianceName, reading.Err)
		failOnError(err, "error preparing message")
		mqttClient.Publish(message)
		log.Infoln("published: ", message)
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	close(done)
	handlers.TurnLEDOff(groveHandler, d3)
	log.Infoln("Closing Grove...")
}

func ledCycle(groveHandler groove.Handler, done chan struct{}) {
loop:
	for {
		select {
		case <-done:
			handlers.TurnLEDOff(groveHandler, d3)
			break loop
		default:
			handlers.TurnLEDOn(groveHandler, d3)
			time.Sleep(500 * time.Millisecond)
			handlers.TurnLEDOff(groveHandler, d3)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

type message struct {
	Source string `json:"name"`
	Time   string `json:"timeStamp"`
	AirQ   int    `json:"readingAQ"`
	Cel    int    `json:"celsius"`
	Far    int    `json:"farenheit"`
	Hum    int    `json:"humidity"`
	Zone   string `json:"timeZone"`
	Error  string `json:"error"`
}

func prepMessage(zone *time.Location, reading int, name string, ipErr error) (string, error) {
	t := time.Now()
	var errMsg string
	if ipErr != nil {
		errMsg = ipErr.Error()
	}
	vals, err := dht.ReadDHT()
	if err != nil {
		errMsg = err.Error()
	}
	msg := &message{Source: name,
		Time:  t.In(zone).Format("20060102150405"),
		AirQ:  reading,
		Cel:   vals.CelsiusTemp,
		Far:   vals.FarenheitTemp,
		Hum:   vals.Humidity,
		Zone:  zone.String(),
		Error: errMsg,
	}

	bdy, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bdy), err
}
