package validator

import "github.com/go-playground/validator/v10"

func Validate(dto interface{}) (map[string]string, bool) {
	dtoValidator := validator.New()
	err := dtoValidator.Struct(dto)
	if err != nil {
		errs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errs[err.Field()] = err.ActualTag()
		}
		return errs, false
	}

	return nil, true
}
