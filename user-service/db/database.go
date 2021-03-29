package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func CreateConnection() (*gorm.DB, error) {
	dbType := "mysql"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "test"
	}
	DBName := os.Getenv("DB_NAME")
	if DBName == "" {
		DBName = "laracom_user"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "test"
	}
	return gorm.Open(
		dbType,
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, 32774, DBName),
	)
}
