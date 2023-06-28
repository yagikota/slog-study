package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func main() {
	// returns zap.Logger, a strongly typed logging API
	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	start := time.Now()

	logger.Info("Hello from zap Logger",
		zap.String("name", "John"),
		zap.Int("age", 9),
		zap.String("email", "john@gmail.com"),
	)

	// convert zap.Logger to zap.SugaredLogger for a more flexible and loose API
	// that's still faster than most other structured logging implementations
	sugar := logger.Sugar()
	sugar.Warnf("something bad is about to happen")
	sugar.Errorw("something bad happened",
		"error", fmt.Errorf("oh no"),
		"answer", 42,
	)

	// you can freely convert back to the base `zap.Logger` type at the boundaries
	// of performance-sensitive operations.
	logger = sugar.Desugar()
	logger.Warn("the operation took longer than expected",
		zap.Int64("time_taken_ms", time.Since(start).Milliseconds()),
	)
}
