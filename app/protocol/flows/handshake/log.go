package handshake

import (
	"github.com/Kalibriumnet/Kalibrium/infrastructure/logger"
	"github.com/Kalibriumnet/Kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
