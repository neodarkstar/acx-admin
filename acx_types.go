package main

// ACXSuiteConfig contains the most common cross platform configurations
type ACXSuiteConfig struct {
	Cassandra CassandraConf
	Kafka     KafkaConfig
	Solr      SolrConfig
}

// CassandraConf basic cassandra configuration
type CassandraConf struct {
	Username          string
	Password          string
	Hosts             []string
	Port              string
	Datacenter        string
	class             string
	ReplicationFactor string
}

// KafkaConfig basic kafka configuration
type KafkaConfig struct {
	Topics            []KafkaTopic
	Group             string
	Hosts             []string
	Port              string
	ReplicationFactor string
}

// KafkaTopic configuration with acx-plus
type KafkaTopic struct {
	Family     string
	Partitions string
	Name       string
}

// SolrConfig basic solr configuration
type SolrConfig struct {
	Hosts []string
	Port  string
}
