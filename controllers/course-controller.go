package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ramirescm/golang-rest-api/models"
	"github.com/ramirescm/golang-rest-api/pkg/course"
)

// Create course
//
//	@Summary	Add a new course
//	@Tags		courses
//	@Accept		json
//	@Produce	json
//	@Param		message	body		models.Course		true	"Course Data"
//	@Router		/courses/ [post]
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var newCourse models.Course
	// json.NewDecoder(r.Body).Decode(&newCourse)
	// err := course.GetService().CreateCourse(&newCourse)
	// if err != nil {
	// 	fmt.Println("Erro GetAllCourses")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
}

// GetCourseById
// @Summary 	Get details of one course
// @Description Get details of one course
// @Tags		courses
// @Accept		text/plain
// @Produce	json
// @Param		message	body		models.Course		true	"Course Data"
//
//	@Router		/courses/{id} [get]
func GetCourseById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro GetCourseById")
		return
	}
	record, err := course.GetService().GetCourseById(id)
	json.NewEncoder(w).Encode(record)
}

// GetAllCourses
//
//	@Summary 		Get details of all courses
//	@Description 	Get details of all courses
//	@Tags 			courses
//	@Accept  		json
//	@Produce  		json
//	@Success 		200 {array} models.Course
//	@Router 		/courses [get]
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	courses, err := course.GetService().GetAllCourses()
	if err != nil {
		fmt.Println("Erro GetAllCourses")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(courses)
	w.WriteHeader(http.StatusOK)
}

// Update a course
//
//	@Summary	Update a course
//	@Tags		courses
//	@Accept		text/plain
//	@Produce	json
//	@Param		message	body		models.Course		true	"Course Data"
//	@Router		/courses/ [put]
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro GetCourseById")
		return
	}
	var updateCourse models.Course
	json.NewDecoder(r.Body).Decode(&updateCourse)
	err = course.GetService().UpdateCourse(id, &updateCourse)
	if err != nil {
		fmt.Println("Falha ao atualizar")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Remove a course
//
//	@Summary	Remove a course
//	@Tags		courses
//	@Accept		text/plain
//	@Produce	json
//	@Param		message	body		models.Course		true	"Course Data"
//	@Router		/courses/ [delete]
func RemoveCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Println("Erro RemoveCourse")
		return
	}
	err = course.GetService().RemoveCourse(id)
	if err != nil {
		fmt.Println("Falha ao remover")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
