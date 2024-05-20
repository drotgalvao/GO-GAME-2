package routes

import (
	"database/sql"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/controllers/users"
	"github.com/gorilla/mux"
)

func UserRouter(dbConn *sql.DB) *mux.Router {
	userRouter := mux.NewRouter().PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		users.CreateUser(w, r, dbConn)
	}).Methods("POST")

	return userRouter
}