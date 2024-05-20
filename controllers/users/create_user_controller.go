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
		utils.HandleError(w, http.StatusBadRequest, "Erro ao decodificar o corpo da requisição: "+err.Error())
		return
	}

	// if userCreationDTO.Name == "" || userCreationDTO.Email == "" || userCreationDTO.Password == "" {
	// 	utils.HandleError(w, http.StatusBadRequest, "Todos os campos devem ser preenchidos.")
	// 	return
	// }
	err = validateUserCreationDTO(userCreationDTO, w)
	if err!= nil {
		return // Se houver um erro de validação, interrompe a execução da função CreateUser
	}

	dbConn, err := db.Connect()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Erro ao conectar ao banco de dados: "+err.Error())
		return
	}
	defer dbConn.Close()

	existingUser, err := repositories.GetUserByEmail(dbConn, userCreationDTO.Email)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Erro ao buscar o usuário: "+err.Error())
		return
	}

	if existingUser != nil {
		utils.HandleError(w, http.StatusConflict, "Email já cadastrado.")
		return
	}

	userResponseDTO, err := repositories.SaveUser(dbConn, userCreationDTO)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Erro ao salvar o usuário: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponseDTO)
}


func validateUserCreationDTO(userCreationDTO models.UserCreationDTO, w http.ResponseWriter) error {
	if userCreationDTO.Name == "" || userCreationDTO.Email == "" || userCreationDTO.Password == "" {
		utils.HandleError(w, http.StatusBadRequest, "Todos os campos devem ser preenchidos.")
	}
	return nil
}