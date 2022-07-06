package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LogPath  string           `yaml:"log-path"`
	DataRoot string           `yaml:"data-root"`
	Accounts []*AccountConfig `yaml:"accounts"`
}

func (c *Config) Write(path string) {
	sink, err := os.Create(path)
	if err != nil {
		log.Error().Err(err).Msg("could not open config to write")
		return
	}
	defer sink.Close()

	encoder := yaml.NewEncoder(sink)
	if err := encoder.Encode(c); err != nil {
		log.Error().Err(err).Msg("could not write config")
	}
}

func (c *Config) ApplyDefaults() {
	for _, ac := range c.Accounts {
		ac.ApplyDefaults()
	}
}

type AccountConfig struct {
	Name       string    `yaml:"name"`
	UUID       uuid.UUID `yaml:"id"`
	Username   string    `yaml:"username"`
	Password   string    `yaml:"password"`
	Server     string    `yaml:"server"`
	SMTPServer string    `yaml:"smtp-server"`
	IMAPServer string    `yaml:"imap-server"`
}

func (ac *AccountConfig) ApplyDefaults() {
	if ac.Server == "gmail" {
		ac.SMTPServer = "smtp.gmail.com:587"
		ac.IMAPServer = "imap.gmail.com:993"
	}

	// Each account must have a UUID to track messages, etc
	// in case the user changes "name"
	if ac.UUID == uuid.Nil {
		ac.UUID = uuid.New()
	}
}

func configPath() string {
	homePath, _ := os.UserHomeDir()
	configPath := filepath.Join(homePath, "amu.config.yaml")
	return configPath
}

func LoadConfig() (*Config, error) {
	configYAML, err := ioutil.ReadFile(configPath())
	if err != nil {
		return nil, errors.Wrapf(err,
			"failed to read config file at %v",
			configPath)
	}
	config := &Config{}
	if err := yaml.Unmarshal(configYAML, config); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml")
	}
	config.ApplyDefaults()
	// write back out
	config.Write(configPath())
	return config, nil
}

func ConfigureLog(conf *Config) {
	var logFile *os.File
	var err error
	if conf.LogPath != "" {
		logFile, err = os.OpenFile(conf.LogPath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		if err != nil {
			panic(err)
		}
	}

	if false {
		log.Logger = log.Output(logFile)
	} // else {
	// 	consoleOut := zerolog.ConsoleWriter{Out: os.Stderr}
	// 	logOut := zerolog.MultiLevelWriter(consoleOut, logFile)
	// 	log.Logger = log.Output(logOut)
	// }

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	// display the file and line number, not the whole path to the file
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.DurationFieldInteger = true
}
