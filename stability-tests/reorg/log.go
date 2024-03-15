package main

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
