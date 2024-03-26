package connmanager

import (
	"github.com/Kalibriumnet/Kalibrium/infrastructure/logger"
	"github.com/Kalibriumnet/Kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
