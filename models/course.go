package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type CourseDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
