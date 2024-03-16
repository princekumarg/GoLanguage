package router

import (
	"github.com/gorilla/mux"
	"github.com/princekumarg/BuildApiDB/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.DeletedAllMovies).Methods("GET")
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("api/movie/{id}", controller.DeletedMovie).Methods("DELETE")
	router.HandleFunc("api/deteteall", controller.DeletedAllMovies).Methods("DELETE")
	return router
}
