package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	Id     int     `json:"course_id"`
	Name   string  `json:"title"`
	Price  float32 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"url"`
}

// fake db
var courses []Course

func (course *Course) IsEmpty() bool {
	return course.Id == 0 && course.Name == ""
}

func main() {
	fmt.Println("Welcome to building apis in golang")
	router := mux.NewRouter()
	router.HandleFunc("/", homeFunction)
	router.HandleFunc("/courses", GetAllCourses)
	router.HandleFunc("/course/{id}",UpdateACourse).Methods("PATCH")
	fmt.Println("starting server...")
	fmt.Println("use CTRL+C to exit")
	http.ListenAndServe(":3000", router)
}

// image more of like, every request will have a request and response
func homeFunction(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method %v", r.Method)
	log.Println("Remote address:", r.RemoteAddr)

	w.Write([]byte("<h1> welcome to api by learn code online </h1>"))

}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all courses getting executed")
	w.Header().Set("Content-Type", "applicaiton/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("Get the courses executed")
	w.Header().Set("Content-Type", "applicaiton/json")

	// using mux for params
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["Id"]) // A to i is ASCII to integer
	if err != nil{
		
		panic("Error while parsing")
	}
	for _, course := range courses {
		if course.Id == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// http.Error(w,"Error in finding the course",404)
	json.NewEncoder(w).Encode("No course found with the given ID")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request){
	payload := r.Body
	if payload == nil{
		json.NewEncoder(w).Encode("Course details not found in request for creating")
	}
	var course Course
	err := json.NewDecoder(payload).Decode(&course)
	if err != nil{
		// w.Write(400)
		json.NewEncoder(w).Encode("Failed to parse the payload")
	}

	course.Id = rand.Int()

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside the json")
	}
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Successfully inserted the course")
}


func UpdateACourse(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method,r.URL.Path)
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // these should be the wild card in request
	// const id, err := 
	log.Println(params)

	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message":"Id is missing in the params"})
		return
	}else {
		id, err := strconv.Atoi(params["id"])
		if err != nil{
			panic("Failed to parse integer")
		}
		for _, course := range(courses) {
			if course.Id == int(id) {
				// updating
				course.Name = params["title"]
				json.NewEncoder(w).Encode(course)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Course not found with id %v",id) })
	}
}