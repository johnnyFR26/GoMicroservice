package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           string `envconfig:"PORT" default:"8080"`
	DatabaseURL    string `envconfig:"DATABASE_URL" required:"true"`
	ExchangeAPIKey string `envconfig:"EXCHANGE_API_KEY" required:"true"`
	RabbitMQURL    string `envconfig:"RABBITMQ_URL" default:"amqp://guest:guest@localhost:5672/"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: nenhum arquivo .env encontrado (continuando com variáveis do sistema)")
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}
	return &cfg
}
