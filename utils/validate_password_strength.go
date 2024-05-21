package utils

import (
	"errors"
	"regexp"
)

func ValidatePasswordStrength(password string) error {
	minLength := 8
	if len(password) < minLength {
		return errors.New("password must have at least 8 characters")
	}

	var hasUpper = regexp.MustCompile("[A-Z]").MatchString
	if !hasUpper(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	var hasLower = regexp.MustCompile("[a-z]").MatchString
	if !hasLower(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	var hasNumber = regexp.MustCompile(`\d`).MatchString
	if !hasNumber(password) {
		return errors.New("password must contain at least one number")
	}

	var hasSpecialChar = regexp.MustCompile(`[@#$%^&*(),.?":{}|<>]`).MatchString
	if !hasSpecialChar(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
