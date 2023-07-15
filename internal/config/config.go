package config

type Config struct {
	HTTP struct {
		IP   string `env:"HTTP-IP" env-default:"localhost"`
		Port int    `env:"HTTP-PORT" env-default:"8002"`
	}
	GRPC struct {
		IP   string `env:"HTTP-IP" env-default:"localhost"`
		Port int    `env:"HTTP-PORT" env-default:"8003"`
	}

	PostgresSQL struct {
		Username string `env:"PG_USER"  env-default:"postgres"`
		Password string `env:"PG_PWD" env-default:"postgres"`
		Host     string `env:"PG_HOST"  env-default:"localhost"`
		Port     string `env:"PG_PORT" env-default:"5432"`
		Database string `env:"PG_DATABASE"  env-default:"restaurant_db"`
	}

	CustomerGRPC struct {
		IP   string `env:"CUSTOMER_IP" env-default:"localhost"`
		Port int    `env:"CUSTOMER_PORT" env-default:"8001"`
	}

	Kafka []string `env:"BROKERS" env-default:"localhost:9092"`
}
