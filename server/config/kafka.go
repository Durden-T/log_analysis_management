package config

import "time"

type Kafka struct {
	Hosts             []string      `mapstructure:"hosts" yaml:"hosts"`
	ReadMinBytes      int           `mapstructure:"read_min_bytes" yaml:"read_min_bytes"`
	ReadMaxBytes      int           `mapstructure:"read_max_bytes" yaml:"read_max_bytes"`
	CommitInterval    time.Duration `mapstructure:"commit_interval" yaml:"commit_interval"`
}
