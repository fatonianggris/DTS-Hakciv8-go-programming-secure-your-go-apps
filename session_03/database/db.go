package database

import (
	"fmt"
	"go-programming-secure-your-go-apps/session_03/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = ""
	port     = 5432
	dbname   = ""
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.User{}, entity.Product{})
}

func GetDB() *gorm.DB {
	return db
}
