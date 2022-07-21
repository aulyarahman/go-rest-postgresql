package handler

import (
	"fmt"
	"github.com/go-chi/render"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"net/http"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method Not Allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource Not Found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad Request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRender(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad Request",
		Message:    err.Error()}
}

func ServerErrorRender(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
