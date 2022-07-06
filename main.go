package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hut8/amu/ent"
	"github.com/rs/zerolog/log"
	_ "github.com/xiaoqidun/entps"
)

func main() {
	ctx := context.Background()
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not find config file at")
		os.Exit(1)
	}
	ConfigureLog(config)
	client, err := makeDB(config.DataRoot)
	if err != nil {
		os.Exit(1)
	}
	engineLog := log.With().Str("component", "engine").Logger()
	engine := NewEngine(
		ctx,
		config,
		client,
		engineLog)
	engine.Run()
}

func makeDB(root string) (*ent.Client, error) {
	dbPath := filepath.Join(root, "mail.db")
	client, err := ent.Open("sqlite3",
		fmt.Sprintf("file:%v", dbPath))
	if err != nil {
		log.Error().Err(err).Str("path", dbPath).
			Msg("failed to open datase")
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed creating schema resources")
		return nil, err
	}
	return client, nil
}
