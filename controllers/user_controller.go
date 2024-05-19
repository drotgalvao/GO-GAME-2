// user_controller.go

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            w.WriteHeader(http.StatusConflict)
            json.NewEncoder(w).Encode(models.ErrorDTO{Code: http.StatusConflict, Message: "Email já cadastrado."})
            return
        } else {
            log.Printf("Erro ao salvar o usuário: %v", err)
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            json.NewEncoder(w).Encode(models.ErrorDTO{Code: http.StatusInternalServerError, Message: err.Error()})
            return
        }
    }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}