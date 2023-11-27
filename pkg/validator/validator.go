package echovalidator

import "github.com/go-playground/validator/v10"

type v struct {
	validator *validator.Validate
}

func New() *v {
	return &v{validator.New()}
}

func (v *v) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
