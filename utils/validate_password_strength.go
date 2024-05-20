package utils

import (
	"errors"
	"regexp"
)

func ValidatePasswordStrength(password string) error {
	minLength := 8
	if len(password) < minLength {
		return errors.New("senha deve ter pelo menos 8 caracteres")
	}

	var hasUpper = regexp.MustCompile("[A-Z]").MatchString
	if !hasUpper(password) {
		return errors.New("senha deve conter pelo menos uma letra maiúscula")
	}

	var hasLower = regexp.MustCompile("[a-z]").MatchString
	if !hasLower(password) {
		return errors.New("senha deve conter pelo menos uma letra minúscula")
	}

	var hasNumber = regexp.MustCompile(`\d`).MatchString
	if !hasNumber(password) {
		return errors.New("senha deve conter pelo menos um número")
	}

	var hasSpecialChar = regexp.MustCompile(`[@#$%^&*(),.?":{}|<>]`).MatchString
	if !hasSpecialChar(password) {
		return errors.New("senha deve conter pelo menos um caractere especial")
	}

	return nil
}