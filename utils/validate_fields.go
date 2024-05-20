package utils

import (
	"errors"
	"reflect"
)

// ValidateFields verifica se todos os campos do DTO estão preenchidos.
func ValidateFields(dto interface{}) error {
	value := reflect.ValueOf(dto)
	if value.Kind()!= reflect.Ptr || value.IsNil() {
		return errors.New("dto must be a pointer type")
	}

	// Obtenha o valor apontado pelo ponteiro
	structValue := value.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		if fieldValue.Kind() == reflect.String && fieldValue.Len() == 0 {
			return errors.New("all fields must be filled")
		}
	}

	return nil
}