// // NOTE: Not well tested, just an illustration of what's possible
package main

import (
	"os"

	"golang.org/x/exp/slog"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"os"

// 	"github.com/fatih/color"
// 	"golang.org/x/exp/slog"
// )

// type PrettyHandlerOptions struct {
// 	SlogOpts slog.HandlerOptions
// }

// type PrettyHandler struct {
// 	slog.Handler
// 	l *log.Logger
// }

// func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
// 	level := r.Level.String() + ":"

// 	switch r.Level {
// 	case slog.LevelDebug:
// 		level = color.MagentaString(level)
// 	case slog.LevelInfo:
// 		level = color.BlueString(level)
// 	case slog.LevelWarn:
// 		level = color.YellowString(level)
// 	case slog.LevelError:
// 		level = color.RedString(level)
// 	}

// 	fields := make(map[string]interface{}, r.NumAttrs())
// 	r.Attrs(func(a slog.Attr) bool {
// 		fields[a.Key] = a.Value.Any()

// 		return true
// 	})

// 	b, err := json.MarshalIndent(fields, "", "  ")
// 	if err != nil {
// 		return err
// 	}

// 	timeStr := r.Time.Format("[15:05:05.000]")
// 	msg := color.CyanString(r.Message)

// 	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

// 	return nil
// }

// func NewPrettyHandler(
// 	out io.Writer,
// 	opts PrettyHandlerOptions,
// ) *PrettyHandler {
// 	h := &PrettyHandler{
// 		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
// 		l:       log.New(out, "", 0),
// 	}

// 	return h
// }

// func main() {
// 	opts := PrettyHandlerOptions{
// 		SlogOpts: slog.HandlerOptions{
// 			Level: slog.LevelDebug,
// 		},
// 	}
// 	handler := NewPrettyHandler(os.Stdout, opts)
// 	logger := slog.New(handler)
// 	logger.Debug(
// 		"executing database query",
// 		slog.String("query", "SELECT * FROM users"),
// 	)
// 	logger.Info("image upload successful", slog.String("image_id", "39ud88"))
// 	logger.Warn(
// 		"storage is 90% full",
// 		slog.String("available_space", "900.1 MB"),
// 	)
// 	logger.Error(
// 		"An error occurred while processing the request",
// 		slog.String("url", "https://example.com"),
// 	)
// }

var appEnv = os.Getenv("APP_ENV")

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if appEnv == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)

	logger.Info("Info message")
}
