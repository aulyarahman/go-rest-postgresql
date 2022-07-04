package handler

import (
	"context"
	"fmt"
	"github.com/aulyarahman/bucketeer/db"
	"github.com/aulyarahman/bucketeer/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

var userIdKey = "userID"

func users(router chi.Router) {
	router.Get("/", getAllUsers)
	router.Post("/", createUsers)
	router.Route("/{userId}", func(r chi.Router) {
		router.Use(UserContext)
		router.Get("/", getUser)
		router.Put("/", updateUser)
		router.Delete("/", deleteUser)
	})
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "userId")
		if itemId == "" {
			render.Render(w, r, ErrorRender(fmt.Errorf("User ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRender(fmt.Errorf("Invalid user ID")))
		}

		ctx := context.WithValue(r.Context(), itemIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func createUsers(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddUser(user); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}

	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(int)
	user, err := dbInstance.GetUserById(userId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRender(err))
		}
		return
	}

	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(int)
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	user, err := dbInstance.UpdateUser(userId, userData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}

	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(int)
	err := dbInstance.DeleteUser(userId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}
}
