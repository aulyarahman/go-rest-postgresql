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

var itemIdKey = "itemID"

func items(router chi.Router) {
	router.Get("/", getAllItems)
	router.Post("/", createItems)
	router.Route("/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/", deleteItem)
	})

}

func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRender(fmt.Errorf("Item iD is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRender(fmt.Errorf("Invalid item ID")))
		}

		ctx := context.WithValue(r.Context(), itemIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func createItems(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddIitem(item); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}

	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIdKey).(int)
	item, err := dbInstance.GetItemById(itemId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRender(err))
		}

		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIdKey).(int)
	itemData := models.Item{}
	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	item, err := dbInstance.UpdateItem(itemId, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIdKey).(int)
	err := dbInstance.DeleteItem(itemId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}
}
