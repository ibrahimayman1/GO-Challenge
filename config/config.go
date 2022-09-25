package config

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// DatabaseConfigurations database configurations
type DatabaseConfigurations struct {
	Dsn string `koanf:"dsn"`
}

// Kafka configurations
type KafkaConfiguration struct {
	Brokers   []string `koanf:"broker"`
	Topic     string   `koanf:"topic"`
	MaxBytes  int      `koanf:"maxbytes"`
	Network   string   `koanf:"network"`
	Partition int      `koanf:"partition"`
}

// Configurations Application wide configurations
type Configurations struct {
	Database DatabaseConfigurations `koanf:"database"`
	Kafka    KafkaConfiguration     `koanf:"kafka"`
}

func LoadConfig() *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yml"), yaml.Parser())
	if err != nil {
		fmt.Printf("Error load configration")
	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {
		fmt.Printf("Error load configration")
	}
	return &configuration
}
