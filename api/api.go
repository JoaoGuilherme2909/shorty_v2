package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/joaoguilherme2909/shorty_v2/store"
	"github.com/joaoguilherme2909/shorty_v2/utils"
)

type postBody struct {
	URL string `json:"url"`
}

func NewHandler(client *store.Client) (http.Handler, error) {

	r := chi.NewMux()

	store := store.Store{
		Client: client,
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/create", handlePost(store))
	r.Get("/{code}", handleGet(store))
	return r, nil
}

func handlePost(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body postBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.JsonResponse(w, http.StatusUnprocessableEntity, map[string]any{
				"Error": "Invalid Body",
			})
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			utils.JsonResponse(w, http.StatusBadRequest, map[string]any{
				"Error": "Invalid URL passed",
			})
			return
		}

		code, err := uuid.NewRandom()

		if err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, map[string]any{
				"Error": "Internal server error",
			})
			return
		}

		err = store.SetUrl(r.Context(), code.String(), body.URL)

		if err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, map[string]any{
				"Error": "Internal server error",
			})
			return
		}

		utils.JsonResponse(w, http.StatusOK, map[string]any{
			"Data": code.String(),
		})
	}
}

func handleGet(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		url, err := store.GetUrl(r.Context(), code)

		if err != nil {
			utils.JsonResponse(w, http.StatusNotFound, map[string]any{
				"Error": "Url not found",
			})
			return
		}

		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
