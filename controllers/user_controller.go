package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/db"
	"github.com/drotgalvao/GO-GAME-2/models"
	"github.com/drotgalvao/GO-GAME-2/repositories"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbConn, err := db.Connect()
	if err!= nil {
		http.Error(w, "Erro ao conectar ao banco de dados.", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	err = repositories.SaveUser(dbConn, newUser)
	if err!= nil {
		log.Printf("Erro ao salvar o usuário: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}