package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/collector"
	"github.com/cjlapao/ms-graph-collector-go/entities"
	"github.com/gorilla/mux"
)

func StartCollectionController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantId := vars["id"]
	if tenantId == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	logger := log.Get()
	execution_context.Get().Authorization.TenantId = tenantId
	collector.RegisterDefaultCollectors()
	collectorService := collector.GetCollectorService()
	err := collectorService.Collect()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.LogError(err)

		error := struct {
			Code    string
			Message string
		}{
			Code:    "400",
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(error)
		return
	}
}

func PostCredentialsController(w http.ResponseWriter, r *http.Request) {
	logger := log.Get()
	var mongodbSvc = mongodb.Get()

	var credential entities.TenantCredentials

	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.LogError(err)

		error := struct {
			Code    string
			Message string
		}{
			Code:    "400",
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(error)
		return
	}

	repo := mongodbSvc.GlobalDatabase().GetCollection("credentials").Repository()
	model, err := mongodb.NewUpdateOneModelBuilder().FilterBy("name", mongodb.Equal, credential.Name).Encode(credential).Build()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.LogError(err)

		error := struct {
			Code    string
			Message string
		}{
			Code:    "400",
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(error)
		return
	}
	_, err = repo.UpsertOne(model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.LogError(err)

		error := struct {
			Code    string
			Message string
		}{
			Code:    "400",
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(error)
		return
	}
}
