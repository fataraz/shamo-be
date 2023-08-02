package http

import (
	"github.com/go-playground/validator/v10"
)

type (
	// Response ...
	Response struct {
		Meta Meta        `json:"meta"`
		Data interface{} `json:"data"`
	}
	// Meta ...
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}
)

// APIResponse : Return
func APIResponse(message string, code int, status string, data interface{}) (resp Response) {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	resp = Response{
		Meta: meta,
		Data: data,
	}
	return
}

// FormatValidationError : Return error
func FormatValidationError(err error) (errors []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return
}
