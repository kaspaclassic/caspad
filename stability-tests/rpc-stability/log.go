package main

import (
	"github.com/casklas/caspad/infrastructure/logger"
	"github.com/casklas/caspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
