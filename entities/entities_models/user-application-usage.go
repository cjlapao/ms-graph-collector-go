package entities_models

type UserApplicationUsage struct {
	ID                      string   `json:"id" bson:"_id"`
	CreatedDateTime         string   `json:"createdDateTime"`
	UserID                  string   `json:"userId"`
	AppID                   string   `json:"appId"`
	IPAddress               string   `json:"ipAddress"`
	ClientAppUsed           string   `json:"clientAppUsed"`
	CorrelationID           string   `json:"correlationId"`
	ConditionalAccessStatus string   `json:"conditionalAccessStatus"`
	IsInteractive           bool     `json:"isInteractive"`
	RiskDetail              string   `json:"riskDetail"`
	RiskLevelAggregated     string   `json:"riskLevelAggregated"`
	RiskLevelDuringSignIn   string   `json:"riskLevelDuringSignIn"`
	RiskState               string   `json:"riskState"`
	RiskEventTypes          []string `json:"riskEventTypes"`
	RiskEventTypesV2        []string `json:"riskEventTypes_v2"`
	ResourceDisplayName     string   `json:"resourceDisplayName"`
	ResourceID              string   `json:"resourceId"`
	Location                Location `json:"location"`
}

type UserApplicationUsageWithDetails struct {
	ID                      string           `json:"id" bson:"_id"`
	CreatedDateTime         string           `json:"createdDateTime"`
	UserID                  string           `json:"userId"`
	AppID                   string           `json:"appId"`
	UserDisplayName         string           `json:"userDisplayName"`
	ApplicationDisplayName  string           `json:"applicationDisplayName"`
	IPAddress               string           `json:"ipAddress"`
	ClientAppUsed           string           `json:"clientAppUsed"`
	CorrelationID           string           `json:"correlationId"`
	ConditionalAccessStatus string           `json:"conditionalAccessStatus"`
	IsInteractive           bool             `json:"isInteractive"`
	RiskDetail              string           `json:"riskDetail"`
	RiskLevelAggregated     string           `json:"riskLevelAggregated"`
	RiskLevelDuringSignIn   string           `json:"riskLevelDuringSignIn"`
	RiskState               string           `json:"riskState"`
	RiskEventTypes          []string         `json:"riskEventTypes"`
	RiskEventTypesV2        []string         `json:"riskEventTypes_v2"`
	ResourceDisplayName     string           `json:"resourceDisplayName"`
	ResourceID              string           `json:"resourceId"`
	Location                Location         `json:"location"`
	User                    UsageUser        `json:"user"`
	Application             UsageApplication `json:"application"`
}

type Location struct {
	City            string `json:"city"`
	State           string `json:"state"`
	CountryOrRegion string `json:"countryOrRegion"`
}

type UsageUser struct {
	ID                string `json:"id" bson:"_id"`
	AccountEnabled    bool   `json:"accountEnabled" bson:"accountEnabled"`
	CreatedDateTime   string `json:"createdDatetime" bson:"createdFatetime"`
	Email             string `json:"email" bson:"email"`
	JobTitle          string `json:"jobTitle" bson:"jobTitle"`
	Surname           string `json:"surname" bson:"surname"`
	GivenName         string `json:"givenName" bson:"givenName"`
	UserPrincipalName string `json:"userPrincipalName" bson:"userPrincipalName"`
	DisplayName       string `json:"displayName" bson:"displayName"`
}

type UsageApplication struct {
	ID               string `json:"id" bson:"_id"`
	DisplayName      string `json:"displayName" bson:"displayName"`
	DeletedDateTime  string `json:"deletedDateTime" bson:"deletedDateTime"`
	Enabled          bool   `json:"accountEnabled" bson:"accountEnabled"`
	ApplicationId    string `json:"appId" bson:"appId"`
	Type             string `json:"type" bson:"type"`
	Recognized       bool   `json:"recognized" bson:"recognized"`
	RecognizedSource string `json:"recognizedSource" bson:"recognizedSource"`
	LastActiveDate   string `json:"lastActiveDate" bson:"lastActiveDate"`
	TotalLoginCount  int64  `json:"totalLoginCount" bson:"totalLoginCount"`
}
