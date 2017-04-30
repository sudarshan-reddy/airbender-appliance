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

	log "github.com/Sirupsen/logrus"
)

const (
	d4           = 4
	a0           = 14
	address      = 0x04
	timeInterval = 3 * time.Millisecond
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func main() {
	cfg, err := loadConfigs()
	failOnError(err, "failed to load configs")

	f, err := os.OpenFile(cfg.LogFileName, os.O_WRONLY|os.O_CREATE, 0755)
	failOnError(err, "failed to create log file")
	defer f.Close()
	log.SetOutput(f)

	groveHandler, err := groove.InitGroove(address)
	failOnError(err, "failed to initialise grove")
	defer groveHandler.Close()

	handlers.TurnLEDOn(groveHandler, d4)
	defer handlers.TurnLEDOff(groveHandler, d4)
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	done := make(chan struct{})

	mqttClient := mq.NewClient(cfg.MQTTTopic, cfg.MQTTURL, cfg.MQTTClient)
	defer mqttClient.Close()

	for reading := range handlers.MonitorAirQuality(done, groveHandler, a0, ticker) {
		if reading.Err != nil {
			log.Fatal(reading.Err)
		}
		message, err := prepMessage(reading.Reading, cfg.ApplianceName)
		if err != nil {
			log.Fatal(err)
		}
		mqttClient.Publish(message)
		log.Infoln("published: ", message)
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	close(done)
	log.Infoln("Closing Grove...")
}

type message struct {
	Source string `json:"name"`
	Time   string `json:"timeStamp"`
	AirQ   int    `json:"readingAQ"`
	Error  string `json:"error"`
}

func prepMessage(reading int, name string) (string, error) {
	t := time.Now()
	msg := &message{Source: name,
		Time: t.Format("20060102150405"),
		AirQ: reading,
	}

	bdy, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bdy), err
}
