package models

import (
	"github.com/go-playground/validator"
	"net/http"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CreatedAt   string `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

func (i *Item) Bind(r *http.Request) error {
	validate := validator.New()
	if err := validate.Struct(i); err != nil {
		return err
	}
	return nil
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
