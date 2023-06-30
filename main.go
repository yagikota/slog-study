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

func addFields(fields map[string]any, a slog.Attr) {
	value := a.Value.Any()
	if _, ok := value.([]slog.Attr); !ok {
		fields[a.Key] = value
		return
	}

	attrs := value.([]slog.Attr)
	// ネストしている場合、再起的にフィールドを探索する。
	innerFields := make(map[string]any, len(attrs))
	for _, attr := range attrs {
		addFields(innerFields, attr)
	}
	fields[a.Key] = innerFields
}

func (h *prettyJSONHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make(map[string]any, r.NumAttrs())
	fields["time"] = r.Time
	fields["level"] = r.Level
	fields["message"] = r.Message

	r.Attrs(func(a slog.Attr) bool {
		addFields(fields, a)
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
		"nest log",
		slog.String("key", "1"),
		slog.Group(
			"inner1",
			slog.String("inner1Key1", "1"),
			slog.String("inner1Key2", "2"),
			slog.Group(
				"inner2",
				slog.String("inner2Key1", "1"),
				slog.String("inner2Key2", "2"),
			),
		),
	)
}
