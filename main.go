package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joaoguilherme2909/shorty_v2/api"
	"github.com/joaoguilherme2909/shorty_v2/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Something went wrong")
		os.Exit(1)
	}
	slog.Info("Application running")
}

func run() error {
	client, err := store.NewClient("localhost:6379", "testredis")

	if err != nil {
		return err
	}

	defer client.Client.Close()

	handler, err := api.NewHandler(client)
	if err != nil {
		return err
	}

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
