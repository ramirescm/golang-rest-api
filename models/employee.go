package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
