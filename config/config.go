package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	server ServerCfg
	dbCfg  DBCfg
)

type DBCfg struct {
	PostgresURI string `envconfig:"POSTGRES_URI" default:"postgresql://localhost:5432"`
}

type ServerCfg struct {
	Port   string `envconfig:"PORT" default:"8000"`
	ApiKey string `envconfig:"API_KEY" default:"123456789"`
}

func InitConfig() {
	configs := []interface{}{
		&server,
		&dbCfg,
	}
	for _, instance := range configs {
		err := envconfig.Process("", instance)
		if err != nil {
			log.Fatalf("unable to init config: %v, err: %v", instance, err)
		}
	}
}

func ServerConfig() ServerCfg {
	return server
}

func DBConfig() DBCfg {
	return dbCfg
}
