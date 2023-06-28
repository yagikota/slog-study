package main

import (
	"golang.org/x/exp/slog"
)

func main() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}
