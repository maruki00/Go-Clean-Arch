package pkg_config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		PG       `yaml:"postgres"`
		RabbitMQ `yaml:"rabbitmq"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"log_level"`
	}

	PG struct {
		PoolMax int    `yaml:"pool_max"`
		DsnURL  string `yaml:"dsn_url"`
	}

	RabbitMQ struct {
		URL string `yaml:"url"`
	}
)

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	conf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer conf.Close()

	decoder := yaml.NewDecoder(conf)

	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
