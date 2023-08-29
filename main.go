package main

import (
	"github.com/happywbfriends/logger/v0"
	"time"
)

func main() {
	logCfg := logger.LogConfig{
		IsPretty: true,
	}

	log := logger.NewLogger(logCfg).With("ver", "1.0")

	log.Infof("App started on %s", time.Now())
}
