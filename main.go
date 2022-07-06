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
	engine := NewEngine(ctx,
		config,
		log.With("component", "engine"))
	engine.Run()
}
