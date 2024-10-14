package rpcclient

import (
	"github.com/ouncenet/ounced/infrastructure/logger"
	"github.com/ouncenet/ounced/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
