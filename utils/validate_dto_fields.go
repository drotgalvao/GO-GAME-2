package utils

import (
	"fmt"
	"reflect"
)

func ValidateDTOFields(dto interface{}) (bool, string) {
	value := reflect.ValueOf(dto)
	if value.Kind()!= reflect.Ptr || value.IsNil() {
		return false, "dto must be a pointer type"
	}

	structValue := value.Elem()

	var errorMessage string
	var allFieldsValid bool = true
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldName := structValue.Type().Field(i).Name
		// lowerCaseFieldName := strings.ToLower(fieldName) // convert field name to lower case comment if needed

		if fieldValue.Kind() == reflect.String && fieldValue.Len() == 0 {
			allFieldsValid = false
			errorMessage = fmt.Sprintf("%s is required", fieldName)
			break
		}
	}

	if!allFieldsValid {
		return false, errorMessage
	}

	return true, ""
}