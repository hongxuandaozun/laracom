package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func CreateConnection() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")

	return gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName),
	)
}
