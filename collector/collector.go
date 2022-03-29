package collector

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/cjlapao/common-go/cache/jwt_token_cache"
	"github.com/cjlapao/ms-graph-collector-go/repositories"
)

var globalCollectorService *CollectorService

type Collector interface {
	Collect() bool
}

type UserCollector interface {
	Collect(userId string) bool
}

type CollectorService struct {
	Collectors []Collector
}

func NewCollectorService() *CollectorService {
	globalCollectorService = &CollectorService{
		Collectors: make([]Collector, 0),
	}

	return globalCollectorService
}

func GetCollectorService() *CollectorService {
	if globalCollectorService != nil {
		return globalCollectorService
	}

	return NewCollectorService()
}

func (p *CollectorService) Register(collectors ...Collector) {
	for _, registerCollector := range collectors {
		found := false
		for _, collector := range p.Collectors {
			if reflect.TypeOf(collector) == reflect.TypeOf(registerCollector) {
				found = true
				break
			}
		}
		if !found {
			p.Collectors = append(p.Collectors, registerCollector)
		}
	}
}

func (c *CollectorService) Collect() error {

	logger.Info("Starting collecting from %v collectors", strconv.Itoa(len(c.Collectors)))
	for _, collector := range c.Collectors {
		name := reflect.TypeOf(collector).Name()
		logger.Command("Collecting from %v", name)
		result := collector.Collect()
		if !result {
			logger.Error("there was an error collecting from " + name)
			// return errors.New("there was an error collecting from " + name)
		}
	}
	logger.Info("Finished processing all collector")

	return nil
}

func (c *CollectorService) GetToken(name string) (string, error) {
	tokenSvc := jwt_token_cache.New()
	token := tokenSvc.Get(name)
	if token == nil || token.IsExpired() {
		repo := repositories.CredentialsRepository{}
		credentials := repo.GetCredential(name)
		if credentials.TenantId == "" {
			return "", errors.New("no credentials found with name " + name)
		}
		endpoint := LoginBaseUrl + "/" + credentials.TenantId + TokenUrl
		loginData := url.Values{}
		loginData.Set("grant_type", "client_credentials")
		loginData.Set("client_id", credentials.ClientId)
		loginData.Set("scope", "https://graph.microsoft.com/.default")
		loginData.Set("client_secret", credentials.ClientSecret)

		client := &http.Client{}
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(loginData.Encode()))
		if err != nil {
			return "", err
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(loginData.Encode())))

		res, err := client.Do(r)
		if err != nil {
			return "", err
		}

		if res.StatusCode != 200 {
			return "", errors.New("bad status code " + res.Status)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return "", err
		}

		var result jwt_token_cache.CachedJwtToken

		err = json.Unmarshal(body, &result)

		if err != nil {
			return "", err
		}

		tokenSvc.Set(name, result)
		return result.AccessToken, nil
	}

	return token.AccessToken, nil
}
