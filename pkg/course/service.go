package course

import (
	"sync"

	"github.com/ramirescm/golang-rest-api/database"
	"github.com/ramirescm/golang-rest-api/models"
	"gorm.io/gorm"
)

type CourseService struct {
	db *gorm.DB
}

var courseService *CourseService
var initOnce sync.Once

func GetService() *CourseService {
	initOnce.Do(initService)
	return courseService
}

func initService() {
	db := database.GetDB()
	courseService = NewCourseService(db)
}

func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{
		db: db,
	}
}

/**
 * Course operations
**/

func (dl *CourseService) GetCourseById(id int64) (models.Course, error) {
	var course models.Course
	result := dl.db.First(&course, id)
	if result.Error != nil {
		return course, result.Error
	}
	return course, nil
}

func (dl *CourseService) GetAllCourses() ([]models.CourseDTO, error) {
	var courses []models.CourseDTO
	result := dl.db.Model(&models.Course{}).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (dl *CourseService) CreateCourse(course *models.Course) error {
	result := dl.db.Create(course)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dl *CourseService) UpdateCourse(id int64, course *models.Course) error {
	dbCourse := models.Course{}
	dbCourse.ID = id
	result := dl.db.Model(&dbCourse).Updates(course)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dl *CourseService) RemoveCourse(id int64) error {
	dbCourse := models.Course{}
	dbCourse.ID = id
	result := dl.db.Delete(&dbCourse)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
