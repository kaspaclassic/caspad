package consensus

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
