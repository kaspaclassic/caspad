package rpc

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
