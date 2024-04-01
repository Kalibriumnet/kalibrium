package prefixmanager

import (
	"github.com/kalibriumnet/kalibrium/infrastructure/logger"
	"github.com/kalibriumnet/kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
