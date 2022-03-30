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

type UserAssignedApplicationCollector struct{}

func (c UserAssignedApplicationCollector) Get(userId string) []msgraph_entities.AppRoleAssignment {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the users assigned applications for %v tenant", tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the token for the tenant %v in the graph api, err %v", tenantId, err.Error())
		return nil
	}

	endpoint := GraphBaseUrl + UsersUrl + "/" + userId + "/appRoleAssignments" + "?$top=" + fmt.Sprintf("%v", TopRecords)
	var result []msgraph_entities.AppRoleAssignment
	client := &http.Client{}
	attempts := constants.RetryCount
	logger.Info("Starting first block of %v users assigned applications", fmt.Sprintf("%v", TopRecords))
	for {
		r, err := http.NewRequest("GET", endpoint, nil)
		logger.Info("Finished all blocks of signins, collected a total of %v", fmt.Sprintf("%v", len(result)))
		if err != nil {
			return nil
		}

		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)
		if err != nil {
			attempts -= 1
			if attempts <= 0 {
				logger.Error("Giving up on the collection of appRoleAssignments due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting appRoleAssignments, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting appRoleAssignments, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
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

		var queryResults []msgraph_entities.AppRoleAssignment
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

	return result
}

func (c UserAssignedApplicationCollector) Collect(userId string, userName string) []msgraph_entities.AppRoleAssignment {
	// Sync users organization apps into our database
	logger := log.Get()
	config := execution_context.Get().Configuration
	logger.Info("Getting assigned applications for user %v", userName)
	if userId != "" {
		usersApps := c.Get(userId)

		if usersApps == nil {
			return nil
		}

		if config.GetBool(constants.DumpToDatabase) {
			usersAppsRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsersApplicationsCollection).Repository()
			for _, userApplication := range usersApps {
				model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, userApplication.ID).Encode(userApplication).Build()
				usersAppsRepo.UpsertOne(model)
			}
			logger.Info("Upserted %v assigned applications into database for user %v", strconv.Itoa(len(usersApps)), userName)
		}

		return usersApps
	} else {
		return nil
	}
}
