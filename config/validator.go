package config

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validate *validator.Validate
}

func NewValidator() *CustomValidator  {
	return &CustomValidator{Validate: validator.New()}
}

func (cv *CustomValidator) TryValidate(entities ...interface{}) error {
	for _, entity := range entities {
		if err := cv.Validate.Struct(entity); err != nil {
			return err
		}
	} 
	return nil
}