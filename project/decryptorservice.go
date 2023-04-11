package main

import (
	"dc.local.decryptorservice/pkg/config"
	"dc.local.decryptorservice/pkg/server"
	"dc.local/zlogger"
	"github.com/pkg/errors"
)

func main() {
	zlogger.SetGlobalLevel(zlogger.DebugLevel)
	logger := zlogger.GetLogger("main")
	logger.Debug().Msg("starting up")

	cnf := new(config.Config)
	err := cnf.New()
	if err != nil {
		e := errors.Wrap(err, "No config found")
		logger.Fatal().Err(e).Msg("")
	}

	server.Start(cnf)
}
