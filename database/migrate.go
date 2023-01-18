package database

import (
	"github.com/ramirescm/golang-rest-api/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&models.Employee{}, &models.Course{})
}
