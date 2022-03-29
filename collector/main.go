package collector

import (
	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/log"
)

var logger = log.Get()
var mongodbSvc = mongodb.Get()
