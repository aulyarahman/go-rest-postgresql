package handler

import (
	"github.com/aulyarahman/bucketeer/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/auth", auth)

	router.Route("/items", items)

	router.Route("/users", users)

	router.Route("/payment", payemnt)

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
