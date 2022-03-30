package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/constants"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type DirectoryObjectCollector struct{}

func (c DirectoryObjectCollector) Get(id string, appName string) *msgraph_entities.ServicePrincipal {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the directoryObject from graph api for app %v in %v tenant", appName, tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the token for the tenant %v in the graph api, err %v", tenantId, err.Error())
		return nil
	}

	var result msgraph_entities.ServicePrincipal
	endpoint := GraphBaseUrl + DirectoryObjectUrl + "/" + id
	client := &http.Client{}
	attempts := constants.RetryCount
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			logger.Info("There was an error creating the request, err %v", err.Error())
			return nil
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			attempts -= 1
			if attempts <= 0 {
				logger.Error("Giving up on the collection of the directoryObject due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting the directoryObject, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting the directoryObject, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logger.Error("There was an error the body of the directoryObject in the response, body is empty or nil, err %v", err.Error())
			break
		}

		err = json.Unmarshal(body, &result)
		if err != nil {
			logger.Error("There was an error when marshaling the result from directoryObject, err %v", err.Error())
			break
		} else {
			logger.Trace("Finished collecting the directoryObject for application %v", appName)
			break
		}
	}

	return &result
}
