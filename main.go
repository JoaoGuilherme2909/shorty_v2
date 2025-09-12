package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/joaoguilherme2909/shorty_v2/api"
	"github.com/joaoguilherme2909/shorty_v2/store/redisStore"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Something went wrong")
	}
	slog.Info("Application running")
}

func run() error {

	connection, err := redisStore.NewRedisClient("localhost:6379", "testredis")

	if err != nil {
		return err
	}

	handler := api.NewHandler(connection)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
