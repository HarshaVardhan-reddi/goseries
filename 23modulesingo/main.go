package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Welcome to go modules")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func Greeter(){
	fmt.Println("Hey hello!!")
}

func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Welcome to golang series on local laptop</h1>"))
}