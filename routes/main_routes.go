package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/drotgalvao/GO-GAME-2/docs"
	"github.com/swaggo/http-swagger"

)

func MainRouter(dbConn *sql.DB) *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) 
	mainRouter.PathPrefix("/user").Handler(UserRouter(dbConn))

	return mainRouter
}