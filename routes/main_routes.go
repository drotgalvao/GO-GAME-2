package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func MainRouter(dbConn *sql.DB) *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.PathPrefix("/user").Handler(UserRouter(dbConn))

	return mainRouter
}