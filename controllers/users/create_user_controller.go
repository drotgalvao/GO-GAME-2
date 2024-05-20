package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	// "github.com/drotgalvao/GO-GAME-2/db"
	"github.com/drotgalvao/GO-GAME-2/models"
	"github.com/drotgalvao/GO-GAME-2/repositories"
	"github.com/drotgalvao/GO-GAME-2/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	var userCreationDTO models.UserCreationDTO
	err := json.NewDecoder(r.Body).Decode(&userCreationDTO)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	done := make(chan bool)

	go func() {
		if err := validateUserCreationDTO(userCreationDTO, w); err != nil {
			utils.HandleError(w, http.StatusBadRequest, err.Error())
		} else {
			processUserCreation(userCreationDTO, w, dbConn)
		}
		done <- true
	}()

	<-done
}

func processUserCreation(userCreationDTO models.UserCreationDTO, w http.ResponseWriter, dbConn *sql.DB) {
	// usar db conn

	existingUser, err := repositories.GetUserByEmail(dbConn, userCreationDTO.Email)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error fetching the user: "+err.Error())
		return
	}

	if existingUser != nil {
		utils.HandleError(w, http.StatusConflict, "email already registered.")
		return
	}

	userResponseDTO, err := repositories.SaveUser(dbConn, userCreationDTO)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error saving the user: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponseDTO)
}

func validateUserCreationDTO(userCreationDTO models.UserCreationDTO, w http.ResponseWriter) error {
	if err := utils.ValidateDTOFields(&userCreationDTO); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return err
	}
	if err := utils.ValidatePasswordStrength(userCreationDTO.Password); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}
