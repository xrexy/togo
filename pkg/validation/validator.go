package validation

import (
	"github.com/go-playground/validator/v10"
)

// ErrorResponse is the error response for validation errors
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
}

// validate is the validator instance
var validate = validator.New()

// ValidateStruct is the validation function
func ValidateStruct(i interface{}) ([]*ErrorResponse, error) {
	var errors []*ErrorResponse
	err := validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors, err
}
