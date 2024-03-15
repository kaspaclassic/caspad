package rpcclient

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
