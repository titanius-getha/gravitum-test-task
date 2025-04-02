package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDsn(name, user, password, host string, port int) string {
	return fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s",
		host,
		user,
		password,
		name,
	)
}

func NewPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	return db, err
}
