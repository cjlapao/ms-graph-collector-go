package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
)

type DirectoryObjectCollector struct{}

func (c DirectoryObjectCollector) Get(id string, appName string) *msgraph_entities.ServicePrincipal {
	collectorSvc := GetCollectorService()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the directoryObject for app %v for %v tenant", appName, tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		return nil
	}

	endpoint := GraphBaseUrl + DirectoryObjectUrl + "/" + id

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
		logger.Error("There was an error getting the directory object for app %v for %v tenant", appName, tenantId)
		return nil
	}

	var result msgraph_entities.ServicePrincipal

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}

	return &result
}
