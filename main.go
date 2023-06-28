package main

import (
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.NewLogLogger(handler, slog.LevelError)

	server := http.Server{
		ErrorLog: logger,
	}
}
