package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/hut8/amu/ent"
	"github.com/rs/zerolog"
)

type Engine struct {
	ctx        context.Context
	config     *Config
	imapClient *client.Client
	db         *ent.Client
	log        zerolog.Logger
}

func NewEngine(
	ctx context.Context,
	config *Config,
	db *ent.Client,
	log zerolog.Logger) *Engine {
	return &Engine{
		ctx:    ctx,
		config: config,
		db:     db,
		log:    log,
	}
}

func (e *Engine) applyConfig() {
	if e.config.Debug {
		basename := fmt.Sprintf("network-debug-%v.log",
			time.Now().UnixNano())
		sinkpath := filepath.Join(e.config.DataRoot, basename)
		sink, err := os.Create(sinkpath)
		if err != nil {
			panic(err)
		}
		e.imapClient.SetDebug(sink)
		e.log.Info().Str("path", sinkpath).
			Msg("imap debug will be written")
	}
}

func (e *Engine) Run() {
	e.applyConfig()
	wg := sync.WaitGroup{}
	for _, ac := range e.config.Accounts {
		e.log.Info().Str("account", ac.Name).
			Msg("running engine on account")
		wg.Add(1)
		go func(ac AccountConfig) {
			e.RunAccount(&ac)
			wg.Done()
		}(*ac)
	}
	wg.Wait()
}

func (e *Engine) RunAccount(ac *AccountConfig) {
	l := e.log.With().Str("account", ac.Name).Logger()
	c, err := client.DialTLS(e.config.Accounts[0].IMAPServer, nil)
	if err != nil {
		l.Error().Err(err).Msg("failed to connect")
		return
	}
	e.imapClient = c
	if err := c.Login(ac.Username, ac.Password); err != nil {
		l.Error().Err(err).
			Msgf("authentication failed for %v",
				ac.Name)
		return
	}

	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	l.Debug().Msg("Mailboxes:")
	for m := range mailboxes {
		l.Debug().Msgf("* %v", m.Name)
	}

	if err := <-done; err != nil {
		l.Error().Err(err).Msg("failed to list mailboxes")
		return
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		l.Error().Err(err).Msg("failed to select inbox")
		return
	}
	l.Debug().Msgf("Flags for INBOX: %v", mbox.Flags)

	l.Info().Msg("exiting engine")
}
