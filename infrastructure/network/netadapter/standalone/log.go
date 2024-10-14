package standalone

import (
	"github.com/ouncenet/ounced/infrastructure/logger"
	"github.com/ouncenet/ounced/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
