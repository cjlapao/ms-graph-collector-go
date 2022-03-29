package startup

import (
	"fmt"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/identity"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/security/encryption"
	"github.com/cjlapao/ms-graph-collector-go/controller"
)

var providers = execution_context.Get()
var MongodbSvc *mongodb.MongoDBService
var logger = log.Get()

func Init() {
	ctx := execution_context.Get().WithDefaultAuthorization()
	ctx.Authorization.Options.TokenDuration = 525600
	ctx.Authorization.ValidationOptions.NotBefore = true
	ctx.Authorization.ValidationOptions.Tenant = false
	MongodbSvc = mongodb.Get()
	ctx.Authorization.WithAudience("carloslapao.com")
	kv := ctx.Authorization.KeyVault
	kv.WithBase64HmacKey("HMAC", providers.Configuration.GetString("JWT_HMAC_PRIVATE_KEY"), encryption.Bit256)
	kv.SetDefaultKey("HMAC")
	identity.Seed(MongodbSvc.GlobalDatabase(), MongodbSvc.GlobalDatabaseName)
	var mongodbSvc = mongodb.Get()
	repo := mongodbSvc.GlobalDatabase().GetCollection("credentials").Repository()
	nCredentials := repo.Pipeline().CountCollection()
	logger.Info("Found %v tenants in the configuration", fmt.Sprintf("%v", nCredentials))
	controller.Init()
}
