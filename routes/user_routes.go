package routes

import (
	"github.com/drotgalvao/GO-GAME-2/controllers/users"
	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	userRouter := mux.NewRouter().PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/create", users.CreateUser).Methods("POST")

	return userRouter
}