package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	logger.Info().
		Str("name", "John").
		Int("age", 9).
		Msg("hello from zerolog")
}
