package router

import (
	"github.com/gorilla/mux"
	"github.com/harshitdixit14/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("UPDATE")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
