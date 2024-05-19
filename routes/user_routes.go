package routes

import (
	"github.com/drotgalvao/GO-GAME-2/controllers"
	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	userRouter := mux.NewRouter().PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/create", controllers.CreateUser).Methods("POST")

	return userRouter
}