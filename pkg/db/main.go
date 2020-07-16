package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Client represents a neo4j client
type Client struct {
	Driver  *neo4j.Driver
	Session *neo4j.Session
	Config  *Config
}

// Config is a neo4j config
type Config struct {
	ServerString string
	Username     string
	Password     string
}

// New creates a new client instance, connected to the database.
func New(cfg *Config) (*Client, error) {
	c := new(Client)
	c.Config = cfg

	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }

	driver, err := neo4j.NewDriver(c.Config.ServerString, neo4j.BasicAuth(c.Config.Username, c.Config.Password, ""), configForNeo4j40)
	if err != nil {
		return nil, err
	}

	c.Driver = &driver

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}
	c.Session = &session

	return c, nil
}
