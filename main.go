package main

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/mux"

	"github.com/drotgalvao/GO-GAME-2/controllers"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)
}
