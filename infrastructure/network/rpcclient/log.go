package rpcclient

import (
	"github.com/Kalibriumnet/Kalibrium/infrastructure/logger"
	"github.com/Kalibriumnet/Kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
