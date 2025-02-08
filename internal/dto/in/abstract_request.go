package dtoin

import "github.com/go-playground/validator"

func Validate(input interface{}) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			return errV
		}
		return err
	}
	return nil
}
