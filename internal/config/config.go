package config

type Config struct {
	HTTP struct {
		IP   string `env:"HTTP_IP"`
		Port int    `env:"HTTP_PORT"`
	}
	GRPC struct {
		IP   string `env:"GRPC_IP"`
		Port int    `env:"GRPC_PORT"`
	}

	PostgresSQL struct {
		Username string `env:"PG_USER"`
		Password string `env:"PG_PWD"`
		Host     string `env:"PG_HOST"`
		Port     string `env:"PG_PORT"`
		Database string `env:"PG_DATABASE"`
	}

	CustomerGRPC struct {
		IP   string `env:"CUSTOMER_IP"`
		Port int    `env:"CUSTOMER_PORT"`
	}

	Kafka []string `env:"KAFKA"`
	Topic string   `env:"TOPIC"`
}
