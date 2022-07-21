package models

import (
	"github.com/go-playground/validator"
	"net/http"
)

type Payment struct {
	ID        int    `json:"id"`
	IdUser    int    `json:"id_user" validate:"required,min=10,max=32"`
	Amount    int    `json:"amount" validate:"required,min=10,max=32"`
	CreatedAt string `json:"created_at"`
}

type ListPayment struct {
	Payments []Payment `json:"payments"`
}

func (i *Payment) Bind(r *http.Request) error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

func (*ListPayment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Payment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
