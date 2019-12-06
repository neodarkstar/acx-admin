package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	c, _ := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "172.22.4.61",
	})

	topic := "ac.init.timeseries.santander50k_auto"

	meta, _ := c.GetMetadata(&topic, true, 5000)

	fmt.Println(len(meta.Topics["ac.init.timeseries.santande50k_auto"].Partitions))
}

func verify(suffix string) bool {

	return true
}
