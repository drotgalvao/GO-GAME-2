package main

import (
	"fmt"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/db"
	"github.com/drotgalvao/GO-GAME-2/routes"
)


// @title GO-GAME-2 API
// @version 1.0
// @description API para o jogo GO-GAME-2
func main() {
	dbConn, err := db.PoolConnection()
	if err != nil {
		fmt.Printf("Erro ao conectar com o banco de dados: %v\n", err)
		return
	}
	routesWithDb := routes.MainRouter(dbConn)
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", routesWithDb)
}
