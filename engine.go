package main

import (
	"context"
	"sync"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/rs/zerolog"
)

type Engine struct {
	ctx        context.Context
	config     *Config
	imapClient *client.Client
	log        zerolog.Logger
}

func NewEngine(
	ctx context.Context,
	config *Config,
	log zerolog.Logger) *Engine {
	return &Engine{
		ctx:    ctx,
		config: config,
		log:    log,
	}
}

func (e *Engine) Run() {
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

	l.Println("Mailboxes:")
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
