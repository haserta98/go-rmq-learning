package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection() *Connection {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable TimeZone=Europe/Istanbul"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Connection{
		DB: db,
	}
}
