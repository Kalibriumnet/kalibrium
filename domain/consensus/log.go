package consensus

import (
	"github.com/kalibriumnet/kalibrium/infrastructure/logger"
	"github.com/kalibriumnet/kalibrium/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
