package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)
	logLevel.Set(slog.LevelError)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logLevel.Set(slog.LevelDebug)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}
