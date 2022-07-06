package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not find config file at")
		os.Exit(1)
	}
	ConfigureLog(config)
	engineLog := log.With().Str("component", "engine").Logger()
	engine := NewEngine(ctx, config, engineLog)
	engine.Run()
}
