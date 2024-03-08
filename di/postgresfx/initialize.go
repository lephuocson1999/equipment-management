package postgresfx

import (
	"equipment-management/config"
	"fmt"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	providePostgresDBClient,
)

func providePostgresDBClient(lifecycle fx.Lifecycle) (*gorm.DB, error) {
	cfg := config.DBConfig()
	// string123 := "host=localhost user=postgres password=example dbname=equipments port=5432"
	db, err := gorm.Open(postgres.Open(cfg.PostgresURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connect to postgres successfully")
	return db, nil
}
