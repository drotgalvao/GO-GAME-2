package main

import (
	"fmt"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/routes"
)



func main() {
	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", routes.MainRouter())
}
