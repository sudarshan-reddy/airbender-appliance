package main

import "github.com/kelseyhightower/envconfig"

// Config returns the local environment variables
type Config struct {
	MQTTURL       string `envconfig:"MQTT_URL" required:"true"`
	MQTTTopic     string `envconfig:"MQTT_TOPIC" required:"true"`
	MQTTClient    string `envconfig:"MQTT_CLIENT" required:"true"`
	ApplianceName string `envconfig:"APPLIANCE_NAME" required:"true"`
	LogFileName   string `envconfig:"LOG_FILE_NAME" required:"true"`
}

func loadConfigs() (*Config, error) {
	var config Config
	err := envconfig.Process("AA", &config)
	return &config, err
}
