package ready

import (
	"github.com/casklas/caspad/infrastructure/logger"
	"github.com/casklas/caspad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
