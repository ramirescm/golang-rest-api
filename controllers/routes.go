package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/api/employees", GetAllEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/api/employees", CreateEmployee).Methods("POST")
	r.HandleFunc("/api/employees/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", RemoveEmployee).Methods("DELETE")

	r.HandleFunc("/api/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", GetCourseById).Methods("GET")
	r.HandleFunc("/api/courses", CreateCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", RemoveCourse).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":9000", r))

}
