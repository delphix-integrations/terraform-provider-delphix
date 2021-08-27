package main

import (
	"log"

	delphix "github.com/ajaytho/delphix-go-sdk"
)

//Config holds our parameters for connecting to the Delphix Engine
type Config struct {
	url      string
	username string
	password string
}

// Client returns a new Datadog client.
func (c *Config) Client() *delphix.Client {

	client := delphix.NewClient(c.username, c.password, c.url)
	log.Printf("[INFO] Delphix Client configured ")

	return client
}
