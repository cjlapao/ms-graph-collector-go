package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/constants"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type SignInCollector struct{}

func (c SignInCollector) Collect(userId string, userName string) []msgraph_entities.SignIn {
	// Sync users signins into our database
	logger := log.Get()
	config := execution_context.Get().Configuration
	signins := c.getByUserId(userId, userName)

	if signins == nil {
		return make([]msgraph_entities.SignIn, 0)
	}

	if config.GetBool(constants.DumpToDatabase) {
		usageRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsageCollection).Repository()

		for _, usage := range signins {
			model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, usage.ID).Encode(usage).Build()
			usageRepo.UpsertOne(model)
		}
		logger.Info("Upserted %v usage signins into database", strconv.Itoa(len(signins)))
	}

	return signins
}

func (c SignInCollector) CollectAll() []msgraph_entities.SignIn {
	// Sync users signins into our database
	logger := log.Get()
	config := execution_context.Get().Configuration
	signins := c.getAll()

	if signins == nil {
		return make([]msgraph_entities.SignIn, 0)
	}

	if config.GetBool(constants.DumpToDatabase) {
		usageRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsageCollection).Repository()

		for _, usage := range signins {
			model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, usage.ID).Encode(usage).Build()
			usageRepo.UpsertOne(model)
		}
		logger.Info("Upserted %v usage signins into database", strconv.Itoa(len(signins)))
	}

	return signins
}

func (c SignInCollector) getByUserId(userId string, userName string) []msgraph_entities.SignIn {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	dateMinus7days := time.Now().Add(((time.Hour * 24) * 7) * -1)
	logger.Info("Starting to collect signins from graph api for user %v in %v tenant from %v", userName, tenantId, dateMinus7days.Format(time.RFC3339))
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the token for the tenant %v in the graph api, err %v", tenantId, err.Error())
		return nil
	}

	endpoint := GraphBaseUrl + SigninsUrl + "?$top=" + fmt.Sprintf("%v", TopRecords) + "&$filter=userId%20eq%20'" + userId + "'%20and%20createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	result := make([]msgraph_entities.SignIn, 0)
	client := &http.Client{}
	attempts := constants.RetryCount
	logger.Info("Starting first block of %v signins", fmt.Sprintf("%v", TopRecords))
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			logger.Info("Finished all blocks of signins, collected a total of %v", fmt.Sprintf("%v", len(result)))
			break
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			attempts -= 1
			if attempts <= 0 {
				logger.Error("Giving up on the collection of signins due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting signins, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting signins, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logger.Error("There was an error the body of the sign in response is empty or nil, err %v", err.Error())
			break
		}

		var rawResult map[string]interface{}
		err = json.Unmarshal(body, &rawResult)
		if err != nil {
			logger.Error("There was an error when unmarshal the body from signins, err %v", err.Error())
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

		var queryResults []msgraph_entities.SignIn
		err = json.Unmarshal(marshaledValue, &queryResults)
		if err != nil {
			logger.Error("There was an error when marshaling the query results from signins, err %v", err.Error())
			break
		}

		result = append(result, queryResults...)

		if rawResult["@odata.nextLink"] != nil && rawResult["@odata.nextLink"].(string) != "" {
			endpoint = rawResult["@odata.nextLink"].(string)
			logger.Info("Starting next block of %v signins, collected %v in a total of %v, sleeping 2 seconds", fmt.Sprintf("%v", TopRecords), fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			time.Sleep(2 * time.Second)
		} else {
			logger.Info("Finished all blocks of signins, collected %v in a total of %v", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			break
		}
	}

	logger.Info("Collected %v signins from graph api for tenant %v", fmt.Sprintf("%v", len(result)), tenantId)
	return result
}

func (c SignInCollector) getAll() []msgraph_entities.SignIn {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	dateMinus7days := time.Now().Add(((time.Hour * 24) * 7) * -1)
	logger.Info("Starting to get the signins usage for %v tenant from %v", tenantId, dateMinus7days.Format(time.RFC3339))
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the token for the tenant %v in the graph api, err %v", tenantId, err.Error())
		return nil
	}

	endpoint := GraphBaseUrl + SigninsUrl + "?$top=" + fmt.Sprintf("%v", TopRecords) + "&$filter=createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	// endpoint := GraphBaseUrl + SigninsUrl + "?$top=999&$filter=userId%20eq%20'cdeb3efb-58ca-4e18-86e3-0300b5a4bf04'%20and%20createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	result := make([]msgraph_entities.SignIn, 0)
	client := &http.Client{}
	attempts := constants.RetryCount
	logger.Info("Starting first block of %v signins", fmt.Sprintf("%v", TopRecords))
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			logger.Info("Finished all blocks of signins, collected a total of %v", fmt.Sprintf("%v", len(result)))
			break
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			attempts -= 1
			if attempts <= 0 {
				logger.Error("Giving up on the collection of signins due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting signins, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting signins, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logger.Error("There was an error the body of the sign in response is empty or nil, err %v", err.Error())
			break
		}

		var rawResult map[string]interface{}
		err = json.Unmarshal(body, &rawResult)
		if err != nil {
			logger.Error("There was an error when unmarshal the body from signins, err %v", err.Error())
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

		var queryResults []msgraph_entities.SignIn
		err = json.Unmarshal(marshaledValue, &queryResults)
		if err != nil {
			logger.Error("There was an error when marshaling the query results from signins, err %v", err.Error())
			break
		}

		result = append(result, queryResults...)

		if rawResult["@odata.nextLink"] != nil && rawResult["@odata.nextLink"].(string) != "" {
			endpoint = rawResult["@odata.nextLink"].(string)
			logger.Info("Starting next block of %v signins, collected %v in a total of %v, sleeping 2 seconds", fmt.Sprintf("%v", TopRecords), fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			time.Sleep(2 * time.Second)
		} else {
			logger.Info("Finished all blocks of signins, collected %v in a total of %v", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			break
		}
	}

	logger.Info("Collected %v signins from graph api for tenant %v", fmt.Sprintf("%v", len(result)), tenantId)
	return result
}
