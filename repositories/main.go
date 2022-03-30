package repositories

import (
	"github.com/cjlapao/common-go/database/mongodb"
)

// Collection
const (
	CredentialsCollectionName = "credentials"
)

var mongodbSvc = mongodb.Get()
