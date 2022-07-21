package models

import (
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
)

type User struct {
	IdUser      int    `json:"id_user"`
	Name        string `json:"name" validate:"required,min=10,max=32"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,max=18"`
	CreatedAt   string `json:"created_at"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (u *User) Bind(r *http.Request) error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
