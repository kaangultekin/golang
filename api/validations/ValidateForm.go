package validations

import (
	"fmt"
	"github.com/go-playground/validator"
	validationStructs "golang/api/structs/validation"
)

func ValidateForm(form interface{}) (vr *validationStructs.ValidationResultStruct) {
	vr = &validationStructs.ValidationResultStruct{
		Success: true,
	}

	validate := validator.New()
	err := validate.Struct(form)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		var errorFields []map[string]string

		for _, fieldError := range validationErrors {
			rule := fieldError.Tag()

			if fieldError.Param() != "" {
				rule = fmt.Sprintf("%s=%s", fieldError.Tag(), fieldError.Param())
			}

			errorFields = append(errorFields, map[string]string{
				"field": fieldError.Field(),
				"rule":  rule,
			})
		}

		vr.Success = false
		vr.ErrorFields = errorFields
	}

	return vr
}
