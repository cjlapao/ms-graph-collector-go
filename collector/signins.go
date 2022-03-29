package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type SignInCollector struct{}

func (c SignInCollector) Get(userId string, userName string) []msgraph_entities.SignIn {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	dateMinus7days := time.Now().Add(((time.Hour * 24) * 7) * -1)
	logger.Info("Starting to get the signins usage for user %v for %v tenant from %v", userName, tenantId, dateMinus7days.Format(time.RFC3339))
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		return nil
	}

	endpoint := GraphBaseUrl + SigninsUrl + "?$top=999&$filter=userId%20eq%20'" + userId + "'%20and%20createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	result := make([]msgraph_entities.SignIn, 0)
	client := &http.Client{}
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			break
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			break
		}

		if res.StatusCode != 200 {
			break
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			break
		}

		var rawResult map[string]interface{}
		err = json.Unmarshal(body, &rawResult)
		if err != nil {
			break
		}

		if rawResult["value"] == nil {
			break
		}

		marshaledValue, err := json.Marshal(rawResult["value"])

		if err != nil {
			break
		}

		var queryResults []msgraph_entities.SignIn
		err = json.Unmarshal(marshaledValue, &queryResults)
		if err != nil {
			break
		}

		result = append(result, queryResults...)

		if rawResult["@odata.nextLink"] != nil && rawResult["@odata.nextLink"].(string) != "" {
			endpoint = rawResult["@odata.nextLink"].(string)
			logger.Info("Starting next block of %v signins", fmt.Sprintf("%v", len(queryResults)))
		} else {
			logger.Info("Finished all blocks of signins")
			break
		}
	}

	logger.Info("Collected %v users signins from api for tenant %v", fmt.Sprintf("%v", len(result)), tenantId)
	return result
}

func (c SignInCollector) GetAll() []msgraph_entities.SignIn {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	dateMinus7days := time.Now().Add(((time.Hour * 24) * 7) * -1)
	logger.Info("Starting to get the signins usage for %v tenant from %v", tenantId, dateMinus7days.Format(time.RFC3339))
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		return nil
	}

	endpoint := GraphBaseUrl + SigninsUrl + "?$top=999&$filter=createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	// endpoint := GraphBaseUrl + SigninsUrl + "?$top=999&$filter=userId%20eq%20'cdeb3efb-58ca-4e18-86e3-0300b5a4bf04'%20and%20createdDateTime%20ge%20" + dateMinus7days.Format(time.RFC3339)
	result := make([]msgraph_entities.SignIn, 0)
	client := &http.Client{}
	attempts := 5
	logger.Info("Starting first block of 999 signins")
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
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
			logger.Error("There was an error the body of the signin response is empty or nil, err %v", err.Error())
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
			logger.Info("Starting next block of 999 signins, collected %v in a total of %v, sleeping 2 seconds", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			time.Sleep(2 * time.Second)
		} else {
			logger.Info("Finished all blocks of signins, collected %v in a total of %v", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			break
		}
	}

	logger.Info("Collected %v signins from api for tenant %v", fmt.Sprintf("%v", len(result)), tenantId)
	return result
}

func (c SignInCollector) Collect(userId string, userName string) []msgraph_entities.SignIn {
	// Sync users signins into our database
	logger := log.Get()
	signins := c.Get(userId, userName)

	if signins == nil {
		return nil
	}

	usageRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsageCollection).Repository()

	for _, usage := range signins {
		model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, usage.ID).Encode(usage).Build()
		usageRepo.UpsertOne(model)
	}
	logger.Info("Upserted %v usage signins into database", strconv.Itoa(len(signins)))
	return signins
}
