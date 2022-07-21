package handler

import (
	"github.com/aulyarahman/bucketeer/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var paymentIdKey = "paymentID"

func payemnt(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Get("/", getAllPayment)
		r.Post("/", createPayment)
	})
}

func getAllPayment(w http.ResponseWriter, r *http.Request) {
	payment, err := dbInstance.GetAllPayment()
	if err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}

	if err := render.Render(w, r, payment); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}

}

func createPayment(w http.ResponseWriter, r *http.Request) {
	pay := &models.Payment{}
	if err := render.Bind(r, pay); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.CreatePayment(pay); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}

	if err := render.Render(w, r, pay); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}
