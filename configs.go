package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config returns the local environment variables
type Config struct {
	MQTTURL         string        `envconfig:"MQTT_URL" required:"true"`
	MQTTTopic       string        `envconfig:"MQTT_TOPIC" required:"true"`
	MQTTClient      string        `envconfig:"MQTT_CLIENT" required:"true"`
	ApplianceName   string        `envconfig:"APPLIANCE_NAME" required:"true"`
	LogFileName     string        `envconfig:"LOG_FILE_NAME" required:"true"`
	TimeZone        string        `envconfig:"TIME_ZONE" required:"true"`
	MonitorInterval time.Duration `envconfig:"MONITOR_INTERVAL" required:"true"`
}

func loadConfigs() (*Config, error) {
	var config Config
	err := envconfig.Process("AA", &config)
	return &config, err
}
