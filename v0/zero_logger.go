package logger

import (
	"github.com/rs/zerolog"
)

type zeroLogger struct {
	logger zerolog.Logger
}

func (l *zeroLogger) With(k, v string) Logger {
	return &zeroLogger{
		logger: l.logger.With().Str(k, v).Logger(),
	}
}

func (l *zeroLogger) Debugf(f string, v ...interface{}) {
	l.logger.Debug().Msgf(f, v...)
}

func (l *zeroLogger) Infof(f string, v ...interface{}) {
	l.logger.Info().Msgf(f, v...)
}

func (l *zeroLogger) Warnf(f string, v ...interface{}) {
	l.logger.Warn().Msgf(f, v...)
}

func (l *zeroLogger) Errorf(f string, v ...interface{}) {
	l.logger.Error().Msgf(f, v...)
}
