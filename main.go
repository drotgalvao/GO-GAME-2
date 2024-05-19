package main

import (
	"fmt"
	"net/http"
	
	"github.com/drotgalvao/GO-GAME-2/db"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, mundo!")
}

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %s\n", err)
		return
	}
	defer dbConn.Close()
	http.HandleFunc("/", helloHandler)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
