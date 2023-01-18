package employee

import (
	"sync"

	"github.com/ramirescm/golang-rest-api/database"
	"github.com/ramirescm/golang-rest-api/models"
	"gorm.io/gorm"
)

type EmployeeService struct {
	db *gorm.DB
}

var employeeService *EmployeeService
var initOnce sync.Once

func GetService() *EmployeeService {
	initOnce.Do(initService)
	return employeeService
}

func initService() {
	db := database.GetDB()
	employeeService = NewEmployeeService(db)
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{
		db: db,
	}
}

/**
 * Employee operations
**/

func (dl *EmployeeService) GetEmployeeById(id int64) (models.Employee, error) {
	var employee models.Employee
	result := dl.db.First(&employee, id)
	if result.Error != nil {
		return employee, result.Error
	}
	return employee, nil

	// return object null if record not found
	// employee := &models.Employee{}
	// result := dl.db.First(employee, id)
	//
	//	if result.Error != nil {
	//		return nil, result.Error
	//	}
	//
	// return employee, nil
}

func (dl *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	result := dl.db.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}

func (dl *EmployeeService) CreateEmployee(employee *models.Employee) error {
	result := dl.db.Create(employee)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dl *EmployeeService) UpdateEmployee(id int64, employee *models.Employee) error {
	dbEmployee := models.Employee{}
	dbEmployee.ID = id
	result := dl.db.Model(&dbEmployee).Updates(employee)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dl *EmployeeService) RemoveEmployee(id int64) error {
	dbEmployee := models.Employee{}
	dbEmployee.ID = id
	result := dl.db.Delete(&dbEmployee)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
