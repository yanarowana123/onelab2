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
		errors.Msg = append(errors.Msg, err.Error())
	}
	return &errors
}
