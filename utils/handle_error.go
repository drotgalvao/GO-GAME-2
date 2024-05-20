package utils

import (
	"encoding/json"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/models"
)

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorDTO{Code: statusCode, Message: message})
}