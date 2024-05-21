package utils

func ValidateSamePassword(password, confirmPassword string) bool {
	return password == confirmPassword
}