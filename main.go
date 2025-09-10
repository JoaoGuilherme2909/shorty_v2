package main

import (
	"log/slog"
	"net/http"
	"time"
)

func main() {

	if err := run(); err != nil {
		slog.Error("Something went wrong")
	}
	slog.Info("Application running")
}

func run() error {
	s := http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
