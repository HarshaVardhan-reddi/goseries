package main

import (
	"fmt"
	"log"
	"mongodbtutorial/router"
	"net/http"
	"os"
)

func main(){
	home, err := os.UserHomeDir(); 
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(home)

	fmt.Println("Welcome to mongo db tutorial in golang and yay its an api")

	router := router.MovieRoutes()

	http.ListenAndServe(":3000",router)
}	