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

	if config.Kafka.ReplicationFactor != "3" {
		t.Error("Invalid Replication Factor")
	}

	if config.Cassandra.ReplicationFactor != "3" {
		t.Error("Invalid Replication Factor")
	}

	if len(config.Kafka.Topics) != 3 {
		t.Error("Invalid Topics")
	}

	if config.Kafka.Topics[0].Family != "metadata" {
		t.Error("Invalid Topic Family")
	}

	if config.Kafka.Topics[0].Partitions != "1" {
		t.Error("Invalid Partition Count")
	}

	if config.Kafka.Group != "busdomain_demo16" {
		t.Error("Invalid Group")
	}
}
