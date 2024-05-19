package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/drotgalvao/GO-GAME-2/controllers"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		response := HealthCheckResponse{
			Status: "OK diego galvao",
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)
}
