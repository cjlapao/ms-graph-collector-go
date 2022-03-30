package dataservices

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cjlapao/common-go/cache/jwt_token_cache"
	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/entities/entities_models"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
	"github.com/cjlapao/ms-graph-collector-go/repositories"
	"github.com/google/uuid"
)

var logger = log.Get()
var mongodbSvc = mongodb.Get()

type Neurons struct{}

func (c *Neurons) ToDataServicesRootObject(user *msgraph_entities.User) *entities_models.DataServiceRoot {
	dataServicesUserRootObject := entities_models.DataServiceRoot{
		User: entities_models.DiscoveryUser{
			Software: entities_models.DiscoveryUserSoftware{
				AssignedApps:              make([]entities_models.DiscoveryUserSoftwareAssignedApp, 0),
				Usage:                     make([]entities_models.DiscoveryUserSoftwareUsage, 0),
				AssignedAppsPkAttrName:    "ApplicationID",
				ApplicationTagsPkAttrName: "ApplicationID",
				UsagePkAttrName:           "ID",
				UsagePkAttrNameID:         "true",
			},
			Department:                 user.Department,
			FirstName:                  user.GivenName,
			LastName:                   user.Surname,
			EmailsPkAttrName:           "Type",
			EmployeeID:                 user.EmployeeID,
			JobTitle:                   user.JobTitle,
			FullName:                   user.DisplayName,
			PhoneNumbersPkAttrName:     "Type",
			EmailsPkAttrNameType:       "true",
			PhoneNumbersPkAttrNameType: "true",
			Application: entities_models.DiscoveryUserApplication{
				AzureAD: entities_models.DiscoveryUserApplicationAzureAD{
					AccountEnabled:         strconv.FormatBool(user.AccountEnabled),
					CreatedDateTime:        user.CreatedDateTime,
					DistinguishedName:      user.OnPremisesDistinguishedName,
					DisplayName:            user.DisplayName,
					ID:                     user.ID,
					LastLogOn:              user.OnPremisesLastSyncDateTime,
					PasswordPolicies:       user.PasswordPolicies,
					ProfileImage:           "",
					SamAccountName:         user.OnPremisesSamAccountName,
					SmartCardLogonRequired: user.PasswordPolicies,
					Status:                 strconv.FormatBool(user.AccountEnabled),
				},
			},
			Emails: make([]entities_models.DiscoveryUserEmail, 0),
			Location: entities_models.DiscoveryUserLocation{
				City:    user.City,
				Country: user.Country,
				Office:  user.OfficeLocation,
				ZipCode: user.PostalCode,
			},
			PhoneNumbers: make([]entities_models.DiscoveryUserPhoneNumbers, 0),
		},
		IdentitiesPkAttrName:       "name",
		Identities_pkAttrName_name: "true",
		Identities:                 make([]entities_models.Identity, 0),
		DiscoveryMetadata: entities_models.DiscoveryMetadata{
			ConnectorsPkAttrName:            "ConnectorId",
			DiscoveryServiceLastUpdateTime:  time.Now().Format(time.RFC3339),
			ProvidersPkAttrName:             "name",
			ProvidersPkAttrNameName:         "true",
			ConnectorsPkAttrNameConnectorID: "true",
			Connectors:                      make([]entities_models.DiscoveryConnector, 0),
			Providers:                       make([]entities_models.DiscoveryProvider, 0),
		},
	}

	generateConnectorObject(&dataServicesUserRootObject)

	generateIdentityObject(user, &dataServicesUserRootObject)

	dataServicesUserRootObject.User.Emails = append(dataServicesUserRootObject.User.Emails, entities_models.DiscoveryUserEmail{
		Email: user.Mail,
		Type:  "Work",
	})

	dataServicesUserRootObject.User.PhoneNumbers = append(dataServicesUserRootObject.User.PhoneNumbers, entities_models.DiscoveryUserPhoneNumbers{
		Number:   user.MobilePhone,
		Type:     "Mobile Device",
		TypeOrig: "Mobile",
	})

	if len(user.BusinessPhones) > 0 {
		dataServicesUserRootObject.User.PhoneNumbers = append(dataServicesUserRootObject.User.PhoneNumbers, entities_models.DiscoveryUserPhoneNumbers{
			Number: user.BusinessPhones[0],
			Type:   "Business",
		})
	}

	return &dataServicesUserRootObject
}

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

func generateConnectorObject(root *entities_models.DataServiceRoot) error {
	azureAdPocConnector := entities_models.DiscoveryConnector{
		ConnectorName:       "Mocker_" + mongodbSvc.TenantDatabaseName,
		ConnectorServerName: "local",
		JobID:               uuid.NewString(),
		Provider:            "AzureADCollector",
		ConnectorID:         "1a2dc6df-9476-4b45-bad9-a10d476edf2f",
		SyncID:              uuid.NewString(),
		DataType:            "User",
		LastConnectorRun:    time.Now().Format(time.RFC3339),
	}
	root.DiscoveryMetadata.Connectors = append(root.DiscoveryMetadata.Connectors, azureAdPocConnector)

	azureAdProvider := entities_models.DiscoveryProvider{
		Name:        "AzureADCollector",
		ProcessDate: azureAdPocConnector.LastConnectorRun,
	}
	root.DiscoveryMetadata.Providers = append(root.DiscoveryMetadata.Providers, azureAdProvider)

	return nil
}

func generateIdentityObject(user *msgraph_entities.User, root *entities_models.DataServiceRoot) error {
	// Email Identity
	emailIdentity := entities_models.Identity{
		Name:  "Email",
		Value: user.Mail,
	}
	// ID Identity
	idIdentity := entities_models.Identity{
		Name:  "AzureADID",
		Value: user.ID,
	}

	// EmailOrDistinguishedName Identity
	emailOrDistinguishedNameIdentity := entities_models.Identity{
		Name:  "Name_EmailOrDistinguishedName",
		Value: user.OnPremisesDistinguishedName,
	}
	root.Identities = append(root.Identities, emailIdentity)
	root.Identities = append(root.Identities, idIdentity)
	root.Identities = append(root.Identities, emailOrDistinguishedNameIdentity)

	return nil
}
