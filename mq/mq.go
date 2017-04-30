package mq

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client is a very light wrapper on mqtt
type Client struct {
	Client mqtt.Client
	Topic  string
}

// NewClient returns a new instance of Client
func NewClient(clientID, raw, topic string) *Client {
	uri, _ := url.Parse(raw)
	server := (fmt.Sprintf("tcp://%s", uri.Host))
	username := uri.User.Username()
	password, _ := uri.User.Password()

	connOpts := mqtt.NewClientOptions().AddBroker(server).SetClientID(clientID).SetCleanSession(true)
	connOpts.SetUsername(username)
	connOpts.SetPassword(password)
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	client := &Client{Client: mqtt.NewClient(connOpts), Topic: topic}

	if token := client.Client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

// Publish publishes a message to the predefined topic
func (m *Client) Publish(message string) error {
	m.Client.Publish(m.Topic, byte(0), false, message)
	return nil
}

// Close disconnects
func (m *Client) Close() {
	m.Client.Disconnect(250)
}
