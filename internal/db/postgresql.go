package db

import (
	configmodels "billing-engine/internal/models/config"
	dbmodels "billing-engine/internal/models/db"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(postgresConfig configmodels.Postgresql) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(
		&dbmodels.User{},
		&dbmodels.Loan{},
		// &dbmodels.Payment{},
		&dbmodels.Billing{},
	)

	return db, nil
}
