package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type UserAssignedApplicationCollector struct{}

func (c UserAssignedApplicationCollector) Get(userId string) []msgraph_entities.AppRoleAssignment {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the users for %v tenant", tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		return nil
	}

	endpoint := GraphBaseUrl + UsersUrl + "/" + userId + "/appRoleAssignments"

	client := &http.Client{}
	r, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil
	}

	r.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(r)
	if err != nil {
		return nil
	}

	if res.StatusCode != 200 {
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil
	}

	var rawResult map[string]interface{}
	var result []msgraph_entities.AppRoleAssignment
	err = json.Unmarshal(body, &rawResult)
	if err != nil {
		return nil
	}

	if rawResult["value"] == nil {
		return nil
	}

	marshaledValue, err := json.Marshal(rawResult["value"])

	if err != nil {
		return nil
	}

	err = json.Unmarshal(marshaledValue, &result)
	if err != nil {
		return nil
	}

	return result
}

func (c UserAssignedApplicationCollector) Collect(userId string, userName string) []msgraph_entities.AppRoleAssignment {
	// Sync users organization apps into our database
	logger := log.Get()
	logger.Info("Getting assigned applications for user %v", userName)
	if userId != "" {
		usersApps := c.Get(userId)

		if usersApps == nil {
			return nil
		}

		usersAppsRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsersOrgApplicationsCollection).Repository()

		for _, user := range usersApps {
			model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, user.ID).Encode(user).Build()
			usersAppsRepo.UpsertOne(model)
		}
		logger.Info("Upserted %v assigned applications applications into database for user %v", strconv.Itoa(len(usersApps)), userName)
		return usersApps
	} else {
		return nil
	}
}
