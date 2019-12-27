package main

import (
	"fmt"

	"github.com/neodarkstar/kafkautil"
	"github.com/spf13/viper"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func loadConfig(path string) (*ACXSuiteConfig, error) {
	viper.Set("Verbose", true)

	var config ACXSuiteConfig

	viper.SetConfigName("acx-suite")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic("Error in Configuration File")
	}
	viper.Unmarshal(&config)

	return &config, err
}

type topicConfig struct {
	Type       string
	Partitions int
}

type config struct {
	Suffix            string
	Prefix            string
	Mode              string
	ReplicationFactor int
	Topics            []topicConfig
}

func readConfigOld() *config {
	// Config Acx Configuration Object
	var acxConfig config

	viper.SetConfigName("acx-suite")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		viper.Set("Suffix", "singus")
		viper.Set("Prefix", "pingus")
		viper.Set("Mode", "mingus")
	}
	viper.Unmarshal(&acxConfig)

	return &acxConfig
}

func main() {
	acxConfig := readConfigOld()

	topics := make([]kafka.TopicSpecification, len(acxConfig.Topics))

	for index, topicConfig := range acxConfig.Topics {
		topic := &kafka.TopicSpecification{
			Topic:             fmt.Sprintf("%s.%s.%s.%s", acxConfig.Prefix, acxConfig.Mode, topicConfig.Type, acxConfig.Suffix),
			NumPartitions:     topicConfig.Partitions,
			ReplicationFactor: acxConfig.ReplicationFactor,
		}

		topics[index] = *topic
	}

	kafkautil.CreateTopics(&topics)

	kafkautil.DeleteTopics(&topics)
}
