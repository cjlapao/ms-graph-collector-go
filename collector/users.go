package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cjlapao/common-go/database/mongodb"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/ms-graph-collector-go/dataservices"
	"github.com/cjlapao/ms-graph-collector-go/entities/entities_models"
	"github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"
	"github.com/google/uuid"
)

type UsersCollector struct{}

func (c UsersCollector) Get() []msgraph_entities.User {
	collectorSvc := GetCollectorService()
	mongodbSvc.TenantDatabase()
	logger := log.Get()
	tenantId := mongodbSvc.TenantDatabaseName
	logger.Info("Starting to get the users for %v tenant", tenantId)
	token, err := collectorSvc.GetToken(tenantId)
	if err != nil {
		return nil
	}

	// endpoint := GraphBaseUrl + UsersUrl + "?$top=500&$filter=userPrincipalName%20eq%20'carlos.lapao@ivanti.com'"
	endpoint := GraphBaseUrl + UsersUrl + ODataTop
	result := make([]msgraph_entities.User, 0)
	client := &http.Client{}
	attempts := 5

	// Getting all of the requests even more that the limit of the api
	logger.Info("Starting first block of 999 users")
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
				logger.Error("Giving up on the collection of users due to exceeding retry attempts")
				break
			} else {
				logger.Error("Retrying because we got an error while collecting users, sleeping for 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if res.StatusCode != 200 {
			attempts -= 1
			if attempts <= 0 {
				break
			} else {
				logger.Error("Retrying because we got an error with status %v while collecting users, sleeping for 5 seconds", fmt.Sprintf("%v", res.StatusCode))
				time.Sleep(5 * time.Second)
				continue
			}
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logger.Error("There was an error the body of the users response is empty or nil, err %v", err.Error())
			break
		}

		var rawResult map[string]interface{}
		err = json.Unmarshal(body, &rawResult)
		if err != nil {
			logger.Error("There was an error when unmarshal the body from users, err %v", err.Error())
			break
		}

		if rawResult["value"] == nil {
			logger.Error("The value property for the users is null, err %v", err.Error())
			break
		}

		marshaledValue, err := json.Marshal(rawResult["value"])

		if err != nil {
			logger.Error("There was an error when marshaling the result from users, err %v", err.Error())
			break
		}

		var queryResults []msgraph_entities.User
		err = json.Unmarshal(marshaledValue, &queryResults)
		if err != nil {
			logger.Error("There was an error when marshaling the query results from users, err %v", err.Error())
			break
		}

		result = append(result, queryResults...)

		if rawResult["@odata.nextLink"] != nil && rawResult["@odata.nextLink"].(string) != "" {
			endpoint = rawResult["@odata.nextLink"].(string)
			logger.Info("Starting next block of 999 users, collected %v in a total of %v, sleeping 2 seconds", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			time.Sleep(2 * time.Second)
		} else {
			logger.Info("Finished all blocks of users, collected %v in a total of %v", fmt.Sprintf("%v", len(queryResults)), fmt.Sprintf("%v", len(result)))
			break
		}
	}

	logger.Info("Collected %v users from api for tenant %v", fmt.Sprintf("%v", len(result)), tenantId)
	return result
}

func (c UsersCollector) Collect() bool {
	// Sync users into our database
	logger.Info("Starting to sync users into the mongodb")
	startCachingTime := time.Now()
	// Getting all of the users
	users := c.Get()

	// Caching objects
	microsoftApps := ServicePrincipalCollector{}.Collect()
	// tenantSignins := SignInCollector{}.GetAll()
	cachedDirectoryObjects := make([]*msgraph_entities.ServicePrincipal, 0)

	if users == nil {
		return false
	}

	dataServiceUsers := make([]entities_models.DataServiceRoot, 0)

	usersRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsersCollection).Repository()

	usersLength := len(users)
	successOp := 0
	errorOp := 0

	batchSizeForPost := 100
	batchedUsers := make([]entities_models.DataServiceRoot, 0)

	logger.Info("Starting to process %v users", strconv.Itoa(usersLength))

	endingCachingTime := time.Since(startCachingTime)
	logger.Info("Took %v seconds to cache entities, cached %v application directory objects and %v microsoft service principals", fmt.Sprintf("%s", endingCachingTime), fmt.Sprintf("%v", len(microsoftApps)), fmt.Sprintf("%v", len(cachedDirectoryObjects)))
	startProcessingTime := time.Now()
	for index, user := range users {
		lgMessage := fmt.Sprintf("Processing %v/%v user %v.", strconv.Itoa(index), strconv.Itoa(usersLength), user.DisplayName)
		logger.Info(lgMessage)
		dataServicesUser := entities_models.DataServiceRoot{
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

		// Connector
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
		dataServicesUser.DiscoveryMetadata.Connectors = append(dataServicesUser.DiscoveryMetadata.Connectors, azureAdPocConnector)

		azureAdProvider := entities_models.DiscoveryProvider{
			Name:        "AzureADCollector",
			ProcessDate: azureAdPocConnector.LastConnectorRun,
		}
		dataServicesUser.DiscoveryMetadata.Providers = append(dataServicesUser.DiscoveryMetadata.Providers, azureAdProvider)

		// Adding Identities
		//Email Identity
		emailIdentity := entities_models.Identity{
			Name:  "Email",
			Value: user.Mail,
		}
		// id Identity
		idIdentity := entities_models.Identity{
			Name:  "AzureADID",
			Value: user.ID,
		}
		// EmailOrDistinguishedName Identity
		emailOrDistinguishedNameIdentity := entities_models.Identity{
			Name:  "Name_EmailOrDistinguishedName",
			Value: user.OnPremisesDistinguishedName,
		}
		dataServicesUser.Identities = append(dataServicesUser.Identities, emailIdentity)
		dataServicesUser.Identities = append(dataServicesUser.Identities, idIdentity)
		dataServicesUser.Identities = append(dataServicesUser.Identities, emailOrDistinguishedNameIdentity)

		dataServicesUser.User.Emails = append(dataServicesUser.User.Emails, entities_models.DiscoveryUserEmail{
			Email: user.Mail,
			Type:  "Work",
		})
		dataServicesUser.User.PhoneNumbers = append(dataServicesUser.User.PhoneNumbers, entities_models.DiscoveryUserPhoneNumbers{
			Number:   user.MobilePhone,
			Type:     "Mobile Device",
			TypeOrig: "Mobile",
		})
		if len(user.BusinessPhones) > 0 {
			dataServicesUser.User.PhoneNumbers = append(dataServicesUser.User.PhoneNumbers, entities_models.DiscoveryUserPhoneNumbers{
				Number: user.BusinessPhones[0],
				Type:   "Business",
			})
		}
		model, _ := mongodb.NewUpdateOneModelBuilder().FilterBy("_id", mongodb.Equal, user.ID).Encode(user).Build()
		_, err := usersRepo.UpsertOne(model)
		if err != nil {
			errorOp += 1
			logger.Error("There was an error upserting the user %v into database", user.DisplayName)
		}

		// Assigned Applications
		userApps := UserAssignedApplicationCollector{}.Collect(user.ID, user.DisplayName)
		for _, app := range userApps {
			var directoryObject *msgraph_entities.ServicePrincipal
			// Querying caching directory object
			for _, cachedDirectoryObject := range cachedDirectoryObjects {
				if cachedDirectoryObject.ID == app.ResourceID {
					logger.Info("Application " + app.ResourceDisplayName + " was found in the cache, processing it.")
					directoryObject = cachedDirectoryObject
				}
			}
			// directoryObject was not found, calling the api and caching it
			if directoryObject == nil {
				logger.Info("Application " + app.ResourceDisplayName + " was not found in the cache, caching it")
				directoryObject = DirectoryObjectCollector{}.Get(app.ResourceID, app.ResourceDisplayName)
				if directoryObject != nil {
					cachedDirectoryObjects = append(cachedDirectoryObjects, directoryObject)
				}
			}

			userApp := entities_models.DiscoveryUserSoftwareAssignedApp{
				ApplicationID:          app.ResourceID,
				AlternateID:            app.ID,
				ApplicationDisplayName: app.ResourceDisplayName,
				PrincipalType:          "User",
				Source:                 "Azure Active Directory",
			}

			if directoryObject != nil && directoryObject.AppID != "" {
				userApp.ApplicationID = directoryObject.AppID
				if len(directoryObject.Tags) > 0 {
					for _, tag := range directoryObject.Tags {
						appTag := entities_models.DiscoveryUserSoftwareApplicationTags{
							ApplicationID: userApp.ApplicationID,
							Tag:           tag,
						}
						dataServicesUser.User.Software.ApplicationTags = append(dataServicesUser.User.Software.ApplicationTags, appTag)
					}
				}
			}

			dataServicesUser.User.Software.AssignedApps = append(dataServicesUser.User.Software.AssignedApps, userApp)
		}

		// Microsoft Internal Applications
		for _, app := range microsoftApps {
			userApp := entities_models.DiscoveryUserSoftwareAssignedApp{
				ApplicationID:          app.AppID,
				AlternateID:            app.AppOwnerOrganizationID,
				ApplicationDisplayName: app.DisplayName,
				PrincipalType:          "Service Principal",
				Source:                 "Azure Active Directory",
			}

			if len(app.Tags) > 0 {
				for _, tag := range app.Tags {
					appTag := entities_models.DiscoveryUserSoftwareApplicationTags{
						ApplicationID: userApp.ApplicationID,
						Tag:           tag,
					}
					dataServicesUser.User.Software.ApplicationTags = append(dataServicesUser.User.Software.ApplicationTags, appTag)
				}
			}

			dataServicesUser.User.Software.AssignedApps = append(dataServicesUser.User.Software.AssignedApps, userApp)
		}

		// Usage
		signins := SignInCollector{}.Collect(user.ID, user.DisplayName)
		for _, application := range dataServicesUser.User.Software.AssignedApps {
			signIn, err := getLastSignInForApplicationUser(signins, application.ApplicationID, user.ID)
			if err == nil {
				userUsage := entities_models.DiscoveryUserSoftwareUsage{
					Result: entities_models.DiscoveryUserSoftwareResult{
						ErrorCode:     "0",
						FailureReason: "None",
						Status:        "success",
					},
					DeviceDetail: entities_models.DiscoveryUserSoftwareDeviceDetail{
						Browser:         signIn.DeviceDetail.Browser,
						IsCompliant:     strconv.FormatBool(signIn.DeviceDetail.IsCompliant),
						IsManaged:       strconv.FormatBool(signIn.DeviceDetail.IsManaged),
						OperatingSystem: signIn.DeviceDetail.OperatingSystem,
					},
					Location: entities_models.DiscoveryUserSoftwareLocation{
						City:            signIn.Location.City,
						CountryOrRegion: signIn.Location.CountryOrRegion,
						State:           signIn.Location.State,
						GeoCoordinates: entities_models.DiscoveryUserSoftwareLocationGeoCoordinates{
							Latitude:  fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Latitude),
							Longitude: fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Longitude),
						},
					},
					AppDisplayName:          signIn.AppDisplayName,
					ApplicationID:           signIn.AppID,
					ConditionalAccessStatus: signIn.ConditionalAccessStatus,
					CorrelationID:           signIn.CorrelationID,
					CreatedDateTime:         signIn.CreatedDateTime,
					ID:                      signIn.ID,
					IPAddress:               signIn.IPAddress,
					IsInteractive:           strconv.FormatBool(signIn.IsInteractive),
					ResourceDisplayName:     signIn.ResourceDisplayName,
					ResourceID:              signIn.AppID,
					RiskDetail:              signIn.RiskDetail,
					RiskLevelDuringSignIn:   signIn.RiskLevelDuringSignIn,
					RiskState:               signIn.RiskState,
					Source:                  "Azure Active Directory",
				}

				logger.Info("Inserted last usage for application %v on %v", signIn.AppDisplayName, signIn.CreatedDateTime)
				dataServicesUser.User.Software.Usage = append(dataServicesUser.User.Software.Usage, userUsage)
			} else {
				// logger.Info("No usage found for application %v and user %v", application.ApplicationDisplayName, user.DisplayName)
			}
		}

		// for _, application := range dataServicesUser.User.Software.AssignedApps {
		// 	signIn, err := getLastSignInForApplicationUser(tenantSignins, application.ApplicationID, user.ID)
		// 	if err == nil {
		// 		userUsage := entities_models.DiscoveryUserSoftwareUsage{
		// 			Result: entities_models.DiscoveryUserSoftwareResult{
		// 				ErrorCode:     "0",
		// 				FailureReason: "None",
		// 				Status:        "success",
		// 			},
		// 			DeviceDetail: entities_models.DiscoveryUserSoftwareDeviceDetail{
		// 				Browser:         signIn.DeviceDetail.Browser,
		// 				IsCompliant:     strconv.FormatBool(signIn.DeviceDetail.IsCompliant),
		// 				IsManaged:       strconv.FormatBool(signIn.DeviceDetail.IsManaged),
		// 				OperatingSystem: signIn.DeviceDetail.OperatingSystem,
		// 			},
		// 			Location: entities_models.DiscoveryUserSoftwareLocation{
		// 				City:            signIn.Location.City,
		// 				CountryOrRegion: signIn.Location.CountryOrRegion,
		// 				State:           signIn.Location.State,
		// 				GeoCoordinates: entities_models.DiscoveryUserSoftwareLocationGeoCoordinates{
		// 					Latitude:  fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Latitude),
		// 					Longitude: fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Longitude),
		// 				},
		// 			},
		// 			AppDisplayName:          signIn.AppDisplayName,
		// 			ApplicationID:           signIn.AppID,
		// 			ConditionalAccessStatus: signIn.ConditionalAccessStatus,
		// 			CorrelationID:           signIn.CorrelationID,
		// 			CreatedDateTime:         signIn.CreatedDateTime,
		// 			ID:                      signIn.ID,
		// 			IPAddress:               signIn.IPAddress,
		// 			IsInteractive:           strconv.FormatBool(signIn.IsInteractive),
		// 			ResourceDisplayName:     signIn.ResourceDisplayName,
		// 			ResourceID:              signIn.AppID,
		// 			RiskDetail:              signIn.RiskDetail,
		// 			RiskLevelDuringSignIn:   signIn.RiskLevelDuringSignIn,
		// 			RiskState:               signIn.RiskState,
		// 			Source:                  "Azure Active Directory",
		// 		}

		// 		logger.Info("Inserted last usage for application %v on %v", signIn.AppDisplayName, signIn.CreatedDateTime)
		// 		dataServicesUser.User.Software.Usage = append(dataServicesUser.User.Software.Usage, userUsage)
		// 	} else {
		// 		// logger.Info("No usage found for application %v and user %v", application.ApplicationDisplayName, user.DisplayName)
		// 	}
		// }
		// signins := SignInCollector{}.Collect(user.ID, user.DisplayName)
		// userSignins := make([]msgraph_entities.SignIn, 0)
		// for _, x := range tenantSignins {
		// 	if x.UserID == user.ID {
		// 		userSignins = append(userSignins, x)
		// 	}
		// }
		// logger.Info("Collected %v user signins from cache", fmt.Sprintf("%v", len(userSignins)))

		// for _, signIn := range userSignins {
		// 	userUsage := entities_models.DiscoveryUserSoftwareUsage{
		// 		Result: entities_models.DiscoveryUserSoftwareResult{
		// 			ErrorCode:     "0",
		// 			FailureReason: "None",
		// 			Status:        "success",
		// 		},
		// 		DeviceDetail: entities_models.DiscoveryUserSoftwareDeviceDetail{
		// 			Browser:         signIn.DeviceDetail.Browser,
		// 			IsCompliant:     strconv.FormatBool(signIn.DeviceDetail.IsCompliant),
		// 			IsManaged:       strconv.FormatBool(signIn.DeviceDetail.IsManaged),
		// 			OperatingSystem: signIn.DeviceDetail.OperatingSystem,
		// 		},
		// 		Location: entities_models.DiscoveryUserSoftwareLocation{
		// 			City:            signIn.Location.City,
		// 			CountryOrRegion: signIn.Location.CountryOrRegion,
		// 			State:           signIn.Location.State,
		// 			GeoCoordinates: entities_models.DiscoveryUserSoftwareLocationGeoCoordinates{
		// 				Latitude:  fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Latitude),
		// 				Longitude: fmt.Sprintf("%f", signIn.Location.GeoCoordinates.Longitude),
		// 			},
		// 		},
		// 		AppDisplayName:          signIn.AppDisplayName,
		// 		ApplicationID:           signIn.AppID,
		// 		ConditionalAccessStatus: signIn.ConditionalAccessStatus,
		// 		CorrelationID:           signIn.CorrelationID,
		// 		CreatedDateTime:         signIn.CreatedDateTime,
		// 		ID:                      signIn.ID,
		// 		IPAddress:               signIn.IPAddress,
		// 		IsInteractive:           strconv.FormatBool(signIn.IsInteractive),
		// 		ResourceDisplayName:     signIn.ResourceDisplayName,
		// 		ResourceID:              signIn.AppID,
		// 		RiskDetail:              signIn.RiskDetail,
		// 		RiskLevelDuringSignIn:   signIn.RiskLevelDuringSignIn,
		// 		RiskState:               signIn.RiskState,
		// 		Source:                  "Azure Active Directory",
		// 	}
		// 	dataServicesUser.User.Software.Usage = append(dataServicesUser.User.Software.Usage, userUsage)
		// }

		dump := execution_context.Get().Configuration.GetBool("dumbData")
		if dump {
			jsonUser, _ := json.MarshalIndent(dataServicesUser, "", "  ")
			userName := strings.ReplaceAll(user.DisplayName, " ", "_")
			filename := ".\\dump\\user" + mongodbSvc.TenantDatabaseName + "_" + userName + ".json"
			if helper.FileExists(filename) {
				helper.DeleteFile(filename)
			}
			helper.WriteToFile(string(jsonUser), filename)
			logger.Success("User %v exported successfully to json on %v", user.DisplayName, filename)
		}

		batchedUsers = append(batchedUsers, dataServicesUser)

		if len(batchedUsers) > batchSizeForPost {
			dataServices := dataservices.Neurons{}
			jsonUserArray, err := json.MarshalIndent(batchedUsers, "", " ")
			if err != nil {
				logger.Error("There was an error converting the batch to the discovery root, err %v", err.Error())
			}

			postAttempts := 5
			for {
				err = dataServices.PostData("Batched Users", string(jsonUserArray))
				if err != nil {
					logger.Error("There was an error posting data to data services we will attempt again in 5 seconds, err %v", err.Error())
					time.Sleep(5 * time.Second)
					postAttempts -= 1
					if postAttempts <= 0 {
						errorOp += 1
						logger.Error("There was an error posting data to data services for more than 5 times, giving up, err %v", err.Error())
						break
					} else {
						continue
					}
				} else {
					logger.Success("Posted %v users successfully to data services", fmt.Sprintf("%v", len(batchedUsers)))
					break
				}
			}
			batchedUsers = make([]entities_models.DataServiceRoot, 0)
		}

		dataServiceUsers = append(dataServiceUsers, dataServicesUser)
		userApplicationsRepo := mongodbSvc.TenantDatabase().GetCollection(MSGraphUsersOrgApplicationsCollection).Repository()
		userApplicationsRepo.FindBy("")
		if errorOp == 0 {
			successOp += 1
		}
	}

	// Posting remaining usage
	if len(batchedUsers) > 0 {
		dataServices := dataservices.Neurons{}
		jsonUserArray, err := json.MarshalIndent(batchedUsers, "", " ")
		if err != nil {
			logger.Error("There was an error converting the batch to the discovery root, err %v", err.Error())
		}

		postAttempts := 5
		for {
			err = dataServices.PostData("Batched Users", string(jsonUserArray))
			if err != nil {
				logger.Error("There was an error posting data to data services we will attempt again in 5 seconds, err %v", err.Error())
				time.Sleep(5 * time.Second)
				postAttempts -= 1
				if postAttempts <= 0 {
					errorOp += 1
					logger.Error("There was an error posting data to data services for more than 5 times, giving up, err %v", err.Error())
					break
				} else {
					continue
				}
			} else {
				logger.Success("Posted %v users successfully to data services", fmt.Sprintf("%v", len(batchedUsers)))
				break
			}
		}
	}

	endingProcessingTime := time.Since(startProcessingTime)
	logger.Info("Took %v seconds to process collector", fmt.Sprintf("%s", endingProcessingTime))
	logger.Info("Processed %v users into database. Success %v | error: %v", strconv.Itoa(usersLength), strconv.Itoa(successOp), strconv.Itoa(errorOp))
	return true
}

func getLastSignInForApplicationUser(signIns []msgraph_entities.SignIn, applicationId string, userId string) (*msgraph_entities.SignIn, error) {

	lastdate, _ := time.Parse(time.RFC3339, "1900-01-01T00:00:00Z")
	var lastSignIn msgraph_entities.SignIn
	for _, signIn := range signIns {
		if strings.EqualFold(signIn.AppID, applicationId) && strings.EqualFold(signIn.UserID, userId) {
			signInDate, err := time.Parse(time.RFC3339, signIn.CreatedDateTime)
			if err != nil {
				logger.Error("There was an error parsing the creation date for the sign in for application %v and user %v", applicationId, userId)
				return nil, err
			}

			if signInDate.After(lastdate) {
				lastdate = signInDate
				lastSignIn = signIn
			}
		}
	}

	validatorDate, _ := time.Parse(time.RFC3339, "1900-01-01T00:00:00Z")
	if lastdate.After(validatorDate) {
		return &lastSignIn, nil
	} else {
		return nil, errors.New("No Usage found")
	}
}
