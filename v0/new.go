package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const defaultLogLevel = zerolog.DebugLevel

type LogConfig struct {
	Level    string `env:"LOG_LEVEL" envDefault:"debug"`
	IsPretty bool   `env:"LOG_PRETTY" envDefault:""`
}

func NewLogger(c LogConfig) Logger {
	if c.IsPretty {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Set correct log level
	if l := c.Level; l != "" {
		level, err := zerolog.ParseLevel(l)
		if err != nil {
			fmt.Printf("invalid log level '%s': %s\n", l, err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Set global log level to %s", level)
		zerolog.SetGlobalLevel(level)
	} else {
		log.Info().Msgf("Global log level set to default: %s", defaultLogLevel)
		zerolog.SetGlobalLevel(defaultLogLevel)
	}

	return &zeroLogger{
		logger: log.Logger,
	}
}
