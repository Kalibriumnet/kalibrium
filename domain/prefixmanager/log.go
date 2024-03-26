package prefixmanager

import (
	"github.com/Kalibriumnet/Kalibrium/infrastructure/logger"
	"github.com/Kalibriumnet/Kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
