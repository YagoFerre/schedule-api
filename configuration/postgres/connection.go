package db

import (
	"fmt"
	"schedule-api/adapters/out/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=123 dbname=schedule_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Printf("PostgreSQL erro: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&entities.Agendamento{})
	if err != nil {
		fmt.Printf("PostgreSQL automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
