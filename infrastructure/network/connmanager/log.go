package connmanager

import (
	"github.com/kalibriumnet/kalibrium/infrastructure/logger"
	"github.com/kalibriumnet/kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
