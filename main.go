package main

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

// slog.Handlerをラップして独自のハンドラーを定義
type prettyJSONHandler struct {
	slog.Handler // Embedded interfaces: https://go.dev/ref/spec#Interface_types
	w            io.Writer
	indent       int
}

func NewJSONPrettyHandler(w io.Writer, opts prettyJSONHandlerOptions) *prettyJSONHandler {
	return &prettyJSONHandler{
		Handler: slog.NewJSONHandler(w, &opts.slogOps),
		w:       w,
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

// Handleメソッドを差し替える
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

	h.w.Write(b)

	return nil
}

func main() {
	ops := prettyJSONHandlerOptions{
		slogOps: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
		indent: 4,
	}
	handler := NewJSONPrettyHandler(os.Stdout, ops)
	logger := slog.New(handler)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	logger.Debug(
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
