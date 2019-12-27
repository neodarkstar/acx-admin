package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := loadConfig(".")

	if err != nil {
		t.Error("Got an Error")
	}

	if config.Cassandra.Username != "cassandra" {
		t.Error("Invalid Username")
	}

	if config.Cassandra.Password != "cassandra" {
		t.Error("Invalid Password")
	}

	if len(config.Cassandra.Hosts) != 1 {
		t.Error("Invalid Hosts")
	}

	if len(config.Solr.Hosts) != 1 {
		t.Error("Invalid Solr Hosts")
	}

	if config.Cassandra.ReplicationFactor != "3" {
		t.Error("Invalid Replication Factor")
	}
}
