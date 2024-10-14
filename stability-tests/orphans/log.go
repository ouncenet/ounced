package main

import (
	"github.com/ouncenet/ounced/infrastructure/logger"
	"github.com/ouncenet/ounced/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)
