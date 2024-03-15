package standalone

import (
	"github.com/kaspaclassic/caspad/infrastructure/logger"
	"github.com/kaspaclassic/caspad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
