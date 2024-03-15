package main

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
