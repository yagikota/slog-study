package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(handler)

	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"incoming request",
		slog.String("method", "GET"),
	)
}
