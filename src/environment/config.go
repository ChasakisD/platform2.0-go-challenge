package environment

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	JWT struct {
		Secret string `yaml:"secret" envconfig:"JWT_SECRET"`
	}
	Server struct {
		Host    string `yaml:"host" envconfig:"SERVER_HOST"`
		Port    string `yaml:"port" envconfig:"SERVER_PORT"`
		Address string
	}
	Database struct {
		Host             string `yaml:"host" envconfig:"DB_HOST"`
		Port             string `yaml:"port" envconfig:"DB_PORT"`
		Username         string `yaml:"username" envconfig:"DB_USER"`
		Password         string `yaml:"password" envconfig:"DB_PASS"`
		Database         string `yaml:"database" envconfig:"DB_NAME"`
		ConnectionString string
	}
}

func LoadConfig(path string) Config {
	config := &Config{}

	yamlFile, err := os.Open(path)
	if err != nil {
		return *config
	}

	_ = yaml.NewDecoder(yamlFile).Decode(config)
	_ = envconfig.Process("", config)

	config.Server.Address = fmt.Sprintf("%s:%s",
		config.Server.Host,
		config.Server.Port)

	config.Database.ConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Database,
		config.Database.Port)

	return *config
}
