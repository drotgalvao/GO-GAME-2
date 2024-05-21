package utils

import (
	"errors"

	"github.com/drotgalvao/GO-GAME-2/models"
)

func ValidateUserCreationDTO(userCreationDTO models.UserCreationDTO) error {
	result, errorMessage := ValidateDTOFields(&userCreationDTO) // check if all fields are setted
	if !result {
		return errors.New(errorMessage)
	}

	if !ValidateSamePassword(userCreationDTO.Password, userCreationDTO.ConfirmPassword) { // check if passwords match
		return errors.New("passwords do not match")
	}

	if err := ValidatePasswordStrength(userCreationDTO.Password); err != nil { // check if password is strong
		return err
	}
	return nil
}
