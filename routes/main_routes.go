package routes

import (
	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.PathPrefix("/user").Handler(UserRouter())

	return mainRouter
}