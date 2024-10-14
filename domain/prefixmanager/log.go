package prefixmanager

import (
	"github.com/ouncenet/ounced/infrastructure/logger"
	"github.com/ouncenet/ounced/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
