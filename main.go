package main

import (
	"context"
	"io"

	"golang.org/x/exp/slog"
)

type prettyJSONHandler struct {
	slog.Handler
	indent int
}

type prettyJSONHandlerOptions struct {
	slogOps slog.HandlerOptions
	indent  int
}

func newPrettyHandler(w io.Writer, opts *prettyJSONHandlerOptions) *prettyJSONHandler {
	return &prettyJSONHandler{
		Handler: slog.NewJSONHandler(w, &opts.slogOps),
		indent:  opts.indent,
	}
}

func (*prettyJSONHandler) Handle(_ context.Context, r slog.Record) error {
 // todo
}

func main() {
	handler := newPrettyHandler()
	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

}
