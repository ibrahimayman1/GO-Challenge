package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type DatabaseConfigurations struct {
	Dsn string `koanf:"dsn"`
}

type KafkaConfiguration struct {
	Brokers   []string `koanf:"broker"`
	Topic     string   `koanf:"topic"`
	MaxBytes  int      `koanf:"maxbytes"`
	Network   string   `koanf:"network"`
	Partition int      `koanf:"partition"`
}

type Configurations struct {
	Database DatabaseConfigurations `koanf:"database"`
	Kafka    KafkaConfiguration     `koanf:"kafka"`
}

func LoadConfig() *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yml"), yaml.Parser())
	if err != nil {

	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {

	}
	return &configuration
}
