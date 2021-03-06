package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/constants"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type ServicePrincipalCollector struct{}

func (c ServicePrincipalCollector) Get() []msgraph_entities.ServicePrincipal {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the Microsoft Service Principals from graph api for %v tenant", tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the token for the tenant %v in the graph api, err %v", tenantId, err.Error())
		return nil
	}

	endpoint := GraphBaseUrl + ServicePrincipalUrl + "?$top=" + fmt.Sprintf("%v", TopRecords)
	var result []msgraph_entities.ServicePrincipal
	client := &http.Client{}
	attempts := 5
	// Getting all of the requests even more that the limit of the api
	logger.Info("Starting first block of %v service principals", fmt.Sprintf("%v", TopRecords))
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			return nil
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			attempts -= 1
			if attempts <= 0 {
				logger.Error("Giving up on the collection of service principals due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting service principals, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting service principals, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logger.Error("There was an error the body of the service principal response is empty or nil, err %v", err.Error())
			break
		}

		var rawResult map[string]interface{}
		err = json.Unmarshal(body, &rawResult)
		if err != nil {
			logger.Error("There was an error when unmarshal the body from service principal, err %v", err.Error())
			break
		}

		if rawResult["value"] == nil {
			logger.Error("The value property for the sigins is null, err %v", err.Error())
			break
		}

		marshaledValue, err := json.Marshal(rawResult["value"])

		if err != nil {
			logger.Error("There was an error when marshaling the result from signins, err %v", err.Error())
			break
		}

		var queryResults []msgraph_entities.ServicePrincipal
		err = json.Unmarshal(marshaledValue, &queryResults)
		if err != nil {
			logger.Error("There was an error when marshaling the query results from service principals, err %v", err.Error())
			break
		}

		result = append(result, queryResults...)

		if rawResult["@odata.nextLink"] != nil && rawResult["@odata.nextLink"].(string) != "" {
			endpoint = rawResult["@odata.nextLink"].(string)
			logger.Info("Starting next block of %v service principals, collected %v in a total of %v, sleeping 2 seconds", fmt.Sprintf("%v", TopRecords), fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			time.Sleep(2 * time.Second)
		} else {
			logger.Info("Finished all blocks of service principals, collected %v in a total of %v", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			break
		}
	}

	return result
}

func (c ServicePrincipalCollector) Collect() []msgraph_entities.ServicePrincipal {
	// Sync service principals into our database
	logger := log.Get()
	microsoftTenantId := constants.MicrosoftTenantId
	config := execution_context.Get().Configuration
	servicePrincipals := c.Get()

	if servicePrincipals == nil {
		return make([]msgraph_entities.ServicePrincipal, 0)
	}

	result := make([]msgraph_entities.ServicePrincipal, 0)

	for _, servicePrincipal := range servicePrincipals {
		if strings.EqualFold(servicePrincipal.AppOwnerOrganizationID, microsoftTenantId) {
			result = append(result, servicePrincipal)
		}
	}

	if config.GetBool(constants.DumpToDatabase) {
		servicePrincipalsRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphServicePrincipalsCollection).Repository()

		for _, msServicePrincipal := range result {
			model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, msServicePrincipal.ID).Encode(msServicePrincipal).Build()
			servicePrincipalsRepo.UpsertOne(model)
		}
		logger.Info("Upserted %v Microsoft Service Principals into database", strconv.Itoa(len(result)))
	}

	return result
}
