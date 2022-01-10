package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/recipes", middleware.GetRecipesHandler).Methods("GET", "OPTIONS")
	return router
}