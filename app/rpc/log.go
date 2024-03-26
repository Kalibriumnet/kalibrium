package rpc

import (
	"github.com/Kalibriumnet/Kalibrium/infrastructure/logger"
	"github.com/Kalibriumnet/Kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
