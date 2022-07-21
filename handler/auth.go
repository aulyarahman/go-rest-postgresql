package handler

import (
	"github.com/aulyarahman/bucketeer/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func auth(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Post("/", sendPhone)
		r.Post("/verify", verifyOtp)
	})
}

func sendPhone(w http.ResponseWriter, r *http.Request) {
	phone := &models.Auth{}
	if err := render.Bind(r, phone); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.Login(phone); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}

	if err := render.Render(w, r, phone); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func verifyOtp(w http.ResponseWriter, r *http.Request) {
	otp := &models.OtpVerify{}
	if err := render.Bind(r, otp); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if _, err := dbInstance.VerifyOTP(otp); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := render.Render(w, r, otp); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}
