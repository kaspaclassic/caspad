package rpc

import (
	"github.com/casklas/caspad/infrastructure/logger"
	"github.com/casklas/caspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
