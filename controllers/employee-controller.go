package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ramirescm/golang-rest-api/models"
	"github.com/ramirescm/golang-rest-api/pkg/employee"
)

// GetEmployeeById godoc
// @Summary      Get an employee
// @Description  get string by ID
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Router       /employees/{id} [get]
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro GetEmployeeById")
		return
	}
	record, err := employee.GetService().GetEmployeeById(id)
	json.NewEncoder(w).Encode(record)

	//const DNS = "root:admin@tcp(127.0.0.1:3308)/golang_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, _ := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	// fmt.Println("teste")
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// var employee models.Employee
	// DB.First(&employee, params["id"])
	// json.NewEncoder(w).Encode(&employee)
}

// GetAllEmployees
//
//	@Summary		Get all employee
//	@Tags 			employees
//	@Router			/employees/ [get]
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := employee.GetService().GetAllEmployees()
	if err != nil {
		fmt.Println("Erro GetAllEmployees")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(employees)
	w.WriteHeader(http.StatusOK)
}

// Create a employee
//
//	@Summary	Create a employee
//	@Tags 		employees
//	@Router		/employees/ [post]
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEmployee models.Employee
	json.NewDecoder(r.Body).Decode(&newEmployee)
	err := employee.GetService().CreateEmployee(&newEmployee)
	if err != nil {
		fmt.Println("Erro GetAllEmployees")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// UpdateEmployee
//
//	@Summary	Update a employee
//	@Tags 		employees
//	@Router		/employees/ [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro GetEmployeeById")
		return
	}
	var updateEmployee models.Employee
	json.NewDecoder(r.Body).Decode(&updateEmployee)
	err = employee.GetService().UpdateEmployee(id, &updateEmployee)
	if err != nil {
		fmt.Println("Falha ao atualizar")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Remove a employee
//
//	@Summary	Remove a employee
//	@Tags 		employees
//	@Router		/employees/ [delete]
func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro RemoveEmployee")
		return
	}
	err = employee.GetService().RemoveEmployee(id)
	if err != nil {
		fmt.Println("Falha ao remover")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
