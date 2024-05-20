package users

import (
	"encoding/json"
	"net/http"

	"github.com/drotgalvao/GO-GAME-2/db"
	"github.com/drotgalvao/GO-GAME-2/models"
	"github.com/drotgalvao/GO-GAME-2/repositories"
	"github.com/drotgalvao/GO-GAME-2/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreationDTO models.UserCreationDTO
	err := json.NewDecoder(r.Body).Decode(&userCreationDTO)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	err = validateUserCreationDTO(userCreationDTO, w)
	if err != nil {
		return
	}

	dbConn, err := db.Connect()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "error connecting to the database: "+err.Error())
		return
	}
	defer dbConn.Close()

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