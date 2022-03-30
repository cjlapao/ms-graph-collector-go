package repositories

import (
	"context"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/ms-graph-collector-go/entities"
	"go.mongodb.org/mongo-driver/bson"
)

var _factory *mongodb.MongoFactory

type CredentialsRepository struct{}

func (c CredentialsRepository) GetCredential(tenantId string) entities.TenantCredentials {
	var result entities.TenantCredentials

	collection := mongodbSvc.GlobalDatabase().GetCollection(CredentialsCollectionName).Repository()

	filter := bson.D{{Key: "tenantId", Value: tenantId}}
	dbResult := collection.FindOne(filter)
	dbResult.Decode(&result)

	return result
}

func (c CredentialsRepository) GetAllCredential() []entities.TenantCredentials {
	result := make([]entities.TenantCredentials, 0)
	ctx := context.Background()
	collection := mongodbSvc.GlobalDatabase().GetCollection(CredentialsCollectionName).Repository()

	filter := bson.D{{}}

	cursor, err := collection.Find(filter)
	if err != nil {
		return nil
	}

	for cursor.Next(ctx) {
		var element entities.TenantCredentials
		cursor.Decode(&element)
		result = append(result, element)
	}

	return result
}
