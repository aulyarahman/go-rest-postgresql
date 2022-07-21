package models

import (
	"github.com/go-playground/validator"
	"net/http"
)

type Auth struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type LoginRequest struct {
	phoneNumber string `json:"phoneNumber"`
}

type OtpVerify struct {
	Id        string `json:"id" validate:"required"`
	Otp       string `json:"otp" validate:"required"`
	CreatedAt string `json:"created_at"`
}

type TapTalkError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Filed   string `json:"filed"`
}

type TapTalkData struct {
	Status      string `json:"status"`
	IsPending   bool   `json:"isPending"`
	IsSent      bool   `json:"isSent"`
	Currency    string `json:"currency"`
	Price       int    `json:"price"`
	CreatedTime int32  `json:"createdTime"`
}

type TapTalkStatus struct {
	Status int32        `json:"status"`
	Error  TapTalkError `json:"error"`
	Data   TapTalkData  `json:"data"`
}

func (i *Auth) Bind(r *http.Request) error {
	validate := validator.New()
	if err := validate.Struct(i); err != nil {
		return err
	}
	return nil
}

func (i *OtpVerify) Bind(r *http.Request) error {
	validate := validator.New()
	if err := validate.Struct(i); err != nil {
		return err
	}
	return nil
}

func (*Auth) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*OtpVerify) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
