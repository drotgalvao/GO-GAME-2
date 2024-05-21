package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

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
		if err := validateUserCreationDTO(userCreationDTO); err != nil {
			utils.HandleError(w, http.StatusBadRequest, err.Error())
		} else {
			processUserCreation(userCreationDTO, w, dbConn)
		}
		done <- true
	}()

	<-done
}

func processUserCreation(userCreationDTO models.UserCreationDTO, w http.ResponseWriter, dbConn *sql.DB) {

	existingUser, err := repositories.GetUserByEmail(dbConn, userCreationDTO.Email)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error fetching the user: "+err.Error())
		return
	}

	if existingUser != nil {
		utils.HandleError(w, http.StatusConflict, "email already registered.")
		return
	}

	hashedPassword, err := utils.HashPassword(userCreationDTO.Password)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error hashing the password: "+err.Error())
		return
	}
	userCreationDTO.Password = hashedPassword

	userResponseDTO, err := repositories.SaveUser(dbConn, userCreationDTO)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error saving the user: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponseDTO)
}

func validateUserCreationDTO(userCreationDTO models.UserCreationDTO) error {
	result, errorMessage := utils.ValidateDTOFields(&userCreationDTO) // check if all fields are setted
	if!result {
		return errors.New(errorMessage)
	}

	if !utils.ValidateSamePassword(userCreationDTO.Password, userCreationDTO.ConfirmPassword) { // check if passwords match
		return errors.New("passwords do not match")
	}

	if err := utils.ValidatePasswordStrength(userCreationDTO.Password); err != nil { // check if password is strong
		return err
	}
	return nil
}
