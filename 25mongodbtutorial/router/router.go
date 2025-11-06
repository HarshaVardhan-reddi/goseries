package router

import (
	"mongodbtutorial/controller"

	"github.com/gorilla/mux"
)


func MovieRoutes() *mux.Router {
	r := mux.NewRouter()	
	r.HandleFunc("/movies",controller.CreateAMovie).Methods("POST")
	r.HandleFunc("/movies/{Id}",controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies",controller.DeleteAllMovie).Methods("DELETE")
	r.HandleFunc("/movies",controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/watched/{id}",controller.GetAllMovies).Methods("PATCH")
	return r
}