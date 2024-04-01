package main

import (
	"github.com/kalibriumnet/kalibrium/infrastructure/logger"
	"github.com/kalibriumnet/kalibrium/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
