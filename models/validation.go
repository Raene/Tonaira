package models

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

//ValidateInputs validates a given dataset against it's struct
func ValidateInputs(dataSet interface{}, val *validator.Validate) (bool, map[string][]string) {
	err := val.Struct(dataSet)

	if err != nil {

		//validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//Validation errors occurred
		errors := make(map[string][]string)
		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "email":
				errors[name] = append(errors[name], "The "+name+" should be a valid email")
				break
			case "eqfield":
				errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
		}

		return false, errors
	}
	return true, nil
}

func InitValidator() *validator.Validate {
	return validator.New()
}
