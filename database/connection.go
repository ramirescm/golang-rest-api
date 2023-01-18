package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func getEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func setupMysql() (*gorm.DB, error) {
	dbHost := getEnvVariable("DB_HOST")
	dbPort := getEnvVariable("DB_PORT")
	dbName := getEnvVariable("DB_DATABASE")
	dbUser := getEnvVariable("DB_USER")
	dbPassword := getEnvVariable("DB_PASSWORD")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeDatabase() error {

	dbs := getEnvVariable("DB")
	var db *gorm.DB
	var err error

	switch dbs {
	case "mysql":
		db, err = setupMysql()
		break
	default:
		return fmt.Errorf("No database found, set the DB env")
	}

	if err != nil {
		return err
	}

	err = Migration(db)
	if err != nil {
		return err
	}
	dbInstance = db
	return nil
}
