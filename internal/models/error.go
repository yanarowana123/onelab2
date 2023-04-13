package models

import "github.com/go-playground/validator/v10"

type ErrorCustom struct {
	Msg string `json:"msg"`
}

type ErrorsCustom struct {
	Msg []string `json:"msg"`
}

func NewErrorsCustomFromValidationErrors(err error) *ErrorsCustom {
	var errors ErrorsCustom
	for _, err := range err.(validator.ValidationErrors) {
		errors.AddError(err.Error())
	}
	return &errors
}

func (m *ErrorsCustom) AddError(msg string) {
	m.Msg = append(m.Msg, msg)
}
