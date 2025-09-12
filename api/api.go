package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/joaoguilherme2909/shorty_v2/store/redisStore"
	"github.com/joaoguilherme2909/shorty_v2/utils"
)

type postBody struct {
	URL string `json:"url"`
}

func NewHandler(connection *redisStore.RedisClient) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/create", handlePost(connection))
	r.Get("/{code}", handleGet(connection))
	return r
}

func handlePost(connection *redisStore.RedisClient) http.HandlerFunc {
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

		err = connection.Client.Set(connection.Ctx, code.String(), body.URL, 1*time.Hour).Err()

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

func handleGet(connection *redisStore.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		url, err := connection.Client.Get(connection.Ctx, code).Result()

		if err != nil {
			utils.JsonResponse(w, http.StatusNotFound, map[string]any{
				"Error": "Url not found",
			})
			return
		}

		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
