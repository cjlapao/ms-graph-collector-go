package dataservices

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/cjlapao/common-go/cache/jwt_token_cache"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/repositories"
)

var logger = log.Get()

type Neurons struct{}

func (c *Neurons) GetToken(name string) (string, error) {
	tokenSvc := jwt_token_cache.New()
	token := tokenSvc.Get("dataservices_" + name)
	if token == nil || token.IsExpired() {
		repo := repositories.CredentialsRepository{}
		credentials := repo.GetCredential(name)
		if credentials.TenantId == "" {
			return "", errors.New("no credentials found with name " + name)
		}

		loginBaseUrl := credentials.UnoLoginUrl
		endpoint := loginBaseUrl + "/" + credentials.NeuronsTenantId + TokenEndpoint
		loginData := url.Values{}
		loginData.Set("grant_type", "client_credentials")
		loginData.Set("client_id", credentials.LoginAppClientId)
		loginData.Set("scope", "web-service")
		loginData.Set("client_secret", credentials.LoginAppClientSecret)

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

func (c *Neurons) PostData(userName string, data string) error {
	logger.Info("Starting to post user data to data services for user %v", userName)
	tenantId := execution_context.Get().Authorization.TenantId
	baseUrl := execution_context.Get().Configuration.GetString("unoBaseUrl")

	if baseUrl == "" {
		repo := repositories.CredentialsRepository{}
		credentials := repo.GetCredential(tenantId)
		baseUrl = credentials.UnoBaseUrl
		if credentials.TenantId == "" {
			return errors.New("no credentials found with name " + tenantId)
		} else {
			execution_context.Get().Configuration.UpsertKey("unoBaseUrl", baseUrl)
		}
	}

	endpoint := baseUrl + DataEndpoint
	client := &http.Client{}

	token, err := c.GetToken(tenantId)
	if err != nil {
		logger.Error("There was an error getting the user token for tenant %v", tenantId)
		return err
	}
	body := strings.NewReader(data)
	r, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		logger.Error("There was an error posting the data to data services for user %v for tenant %v\nerr:", userName, tenantId, err.Error())
		return err
	}

	if res.StatusCode != 200 {
		logger.Error("There was an error posting the data to data services for user %v for tenant %v\nerr: status code %v", userName, tenantId, strconv.Itoa(res.StatusCode))
		return errors.New("wrong code")
	}

	logger.Success("Successfully posted data to data services for user %v for tenant %v", userName, tenantId)
	return nil
}
