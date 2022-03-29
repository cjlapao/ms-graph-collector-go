package msgraph_entities

type SignIn struct {
	ID                               string        `json:"id" bson:"_id"`
	CreatedDateTime                  string        `json:"createdDateTime"`
	UserDisplayName                  string        `json:"userDisplayName"`
	UserPrincipalName                string        `json:"userPrincipalName"`
	UserID                           string        `json:"userId"`
	AppID                            string        `json:"appId"`
	AppDisplayName                   string        `json:"appDisplayName"`
	IPAddress                        string        `json:"ipAddress"`
	ClientAppUsed                    string        `json:"clientAppUsed"`
	CorrelationID                    string        `json:"correlationId"`
	ConditionalAccessStatus          string        `json:"conditionalAccessStatus"`
	IsInteractive                    bool          `json:"isInteractive"`
	RiskDetail                       string        `json:"riskDetail"`
	RiskLevelAggregated              string        `json:"riskLevelAggregated"`
	RiskLevelDuringSignIn            string        `json:"riskLevelDuringSignIn"`
	RiskState                        string        `json:"riskState"`
	RiskEventTypes                   []interface{} `json:"riskEventTypes"`
	RiskEventTypesV2                 []interface{} `json:"riskEventTypes_v2"`
	ResourceDisplayName              string        `json:"resourceDisplayName"`
	ResourceID                       string        `json:"resourceId"`
	Status                           Status        `json:"status"`
	DeviceDetail                     DeviceDetail  `json:"deviceDetail"`
	Location                         Location      `json:"location"`
	AppliedConditionalAccessPolicies []interface{} `json:"appliedConditionalAccessPolicies"`
}
