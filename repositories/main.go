package repositories

import (
	"github.com/cjlapao/common-go/database/mongodb"
)

// Collection
const (
	UsersCollectionName       = "Users"
	CampainsCollectionName    = "Campaigns"
	CredentialsCollectionName = "Credentials"
)

var mongodbSvc = mongodb.Get()
