package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           string `envconfig:"PORT" default:"8080"`
	DatabaseURL    string `envconfig:"DATABASE_URL" required:"true"`
	RabbitMQURL    string `envconfig:"RABBITMQ_URL" default:"amqp://guest:guest@rabbitmq:5672/"`
	ExchangeAPIKey string `envconfig:"EXCHANGE_API_KEY" required:"true"`
}

func Load() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}
	return &cfg
}
