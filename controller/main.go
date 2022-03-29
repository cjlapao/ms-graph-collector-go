package controller

import (
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/identity/database"
	"github.com/cjlapao/common-go/restapi"
)

var listener *restapi.HttpListener
var svc = execution_context.Get()

func Init() {
	listener = restapi.GetHttpListener()
	listener.AddJsonContent().AddLogger().AddHealthCheck().WithAuthentication(database.MongoDBUserContextAdapter{})

	// MS Graph Endpoints
	listener.AddController(StartCollectionController, "/{id}/start", "GET")
	listener.AddController(PostCredentialsController, "/credentials", "POST")

	listener.Start()
}
