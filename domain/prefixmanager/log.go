package prefixmanager

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
