package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

// slog.Handlerをラップして独自のハンドラーを定義
type prettyJSONHandler struct {
	slog.Handler
	l      *log.Logger
	indent int
}

func newPrettyHandler(w io.Writer, opts prettyJSONHandlerOptions) *prettyJSONHandler {
	return &prettyJSONHandler{
		Handler: slog.NewJSONHandler(w, &opts.slogOps),
		l:       log.New(w, "", 0),
		indent:  opts.indent,
	}
}

// 独自のハンドラーに対するオプション
type prettyJSONHandlerOptions struct {
	slogOps slog.HandlerOptions
	indent  int
}

func makeFields(fields map[string]interface{}, a slog.Attr) {
	value := a.Value.Any()
	if _, ok := value.([]slog.Attr); !ok {
		fields[a.Key] = value
		return
	}

	innerFields := make(map[string]interface{}, len(value.([]slog.Attr)))
	for _, attr := range value.([]slog.Attr) {
		makeFields(innerFields, attr)
	}
	fields[a.Key] = innerFields
}

func (h *prettyJSONHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make(map[string]interface{}, r.NumAttrs())
	fields["time"] = r.Time
	fields["level"] = r.Level
	fields["message"] = r.Message
	r.Attrs(func(a slog.Attr) bool {
		makeFields(fields, a)
		return true
	})

	b, err := json.MarshalIndent(fields, "", strings.Repeat(" ", h.indent))
	if err != nil {
		return err
	}

	h.l.Println(string(b))

	return nil
}

func main() {
	ops := prettyJSONHandlerOptions{
		slogOps: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
		indent: 4,
	}
	handler := newPrettyHandler(os.Stdout, ops)
	logger := slog.New(handler)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logger.Info(
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Group("properties",
				slog.Int("height", 3000),
				slog.String("format", "jpeg")),
		),
	)
}
