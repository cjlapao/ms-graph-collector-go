package entities_models

type DataServiceRoot struct {
	IdentitiesPkAttrName       string            `json:"Identities_pkAttrName"`
	Identities_pkAttrName_name string            `json:"Identities_pkAttrName_name"`
	DiscoveryMetadata          DiscoveryMetadata `json:"DiscoveryMetadata"`
	Identities                 []Identity        `json:"Identities"`
	User                       DiscoveryUser     `json:"User"`
}

type DiscoveryMetadata struct {
	ConnectorsPkAttrName            string               `json:"Connectors_pkAttrName"`
	DiscoveryServiceLastUpdateTime  string               `json:"DiscoveryServiceLastUpdateTime"`
	ProvidersPkAttrName             string               `json:"Providers_pkAttrName"`
	ConnectorsPkAttrNameConnectorID string               `json:"Connectors_pkAttrName_ConnectorId"`
	ProvidersPkAttrNameName         string               `json:"Providers_pkAttrName_name"`
	Connectors                      []DiscoveryConnector `json:"Connectors"`
	Providers                       []DiscoveryProvider  `json:"Providers"`
}

type DiscoveryConnector struct {
	ConnectorName       string `json:"ConnectorName"`
	ConnectorServerName string `json:"ConnectorServerName"`
	JobID               string `json:"JobID"`
	Provider            string `json:"Provider"`
	ConnectorID         string `json:"ConnectorId"`
	SyncID              string `json:"SyncId"`
	DataType            string `json:"DataType"`
	LastConnectorRun    string `json:"LastConnectorRun"`
}

type DiscoveryProvider struct {
	Name        string `json:"name"`
	ProcessDate string `json:"processDate"`
}

type Identity struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DiscoveryUser struct {
	Application                DiscoveryUserApplication    `json:"Application"`
	Department                 string                      `json:"Department"`
	EmployeeID                 string                      `json:"EmployeeID"`
	FirstName                  string                      `json:"FirstName"`
	LastName                   string                      `json:"LastName"`
	FullName                   string                      `json:"FullName"`
	Manager                    string                      `json:"Manager"`
	EmailsPkAttrName           string                      `json:"Emails_pkAttrName"`
	EmailsPkAttrNameType       string                      `json:"Emails_pkAttrName_Type"`
	PhoneNumbersPkAttrName     string                      `json:"PhoneNumbers_pkAttrName"`
	PhoneNumbersPkAttrNameType string                      `json:"PhoneNumbers_pkAttrName_Type"`
	Emails                     []DiscoveryUserEmail        `json:"Emails"`
	Software                   DiscoveryUserSoftware       `json:"Software"`
	JobTitle                   string                      `json:"JobTitle"`
	Location                   DiscoveryUserLocation       `json:"Location"`
	PhoneNumbers               []DiscoveryUserPhoneNumbers `json:"PhoneNumbers"`
}

// Generated by https://quicktype.io

type DiscoveryUserPhoneNumbers struct {
	Number   string `json:"Number"`
	Type     string `json:"Type"`
	TypeOrig string `json:"Type_orig,omitempty"`
}

type DiscoveryUserApplication struct {
	AzureAD DiscoveryUserApplicationAzureAD `json:"AzureAD"`
}

type DiscoveryUserEmail struct {
	Email string `json:"Email"`
	Type  string `json:"Type"` // Generated by https://quicktype.io
}

type DiscoveryUserApplicationAzureAD struct {
	AccountEnabled         string `json:"AccountEnabled"`
	CreatedDateTime        string `json:"CreatedDateTime"`
	DistinguishedName      string `json:"DistinguishedName"`
	DisplayName            string `json:"DisplayName"`
	ID                     string `json:"ID"`
	LastLogOn              string `json:"LastLogOn"`
	PasswordPolicies       string `json:"PasswordPolicies"`
	ProfileImage           string `json:"ProfileImage"`
	SamAccountName         string `json:"SamAccountName"`
	SmartCardLogonRequired string `json:"SmartCardLogonRequired"`
	Status                 string `json:"Status"`
}

type DiscoveryUserSoftware struct {
	Usage                                  []DiscoveryUserSoftwareUsage           `json:"Usage"`
	ApplicationTags                        []DiscoveryUserSoftwareApplicationTags `json:"ApplicationTags"`
	AssignedApps                           []DiscoveryUserSoftwareAssignedApp     `json:"AssignedApps"`
	UsagePkAttrName                        string                                 `json:"Usage_pkAttrName"`
	UsagePkAttrNameID                      string                                 `json:"Usage_pkAttrName_ID"`
	AssignedAppsPkAttrName                 string                                 `json:"AssignedApps_pkAttrName"`
	AssignedAppsPkAttrNameApplicationID    string                                 `json:"AssignedApps_pkAttrName_ApplicationID"`
	ApplicationTagsPkAttrName              string                                 `json:"ApplicationTags_pkAttrName"`
	ApplicationTagsPkAttrNameApplicationID string                                 `json:"ApplicationTags_pkAttrName_ApplicationID"`
}

type DiscoveryUserSoftwareUsage struct {
	Result                  DiscoveryUserSoftwareResult       `json:"Result"`
	DeviceDetail            DiscoveryUserSoftwareDeviceDetail `json:"DeviceDetail"`
	Location                DiscoveryUserSoftwareLocation     `json:"Location"`
	ApplicationID           string                            `json:"ApplicationID"`
	AppDisplayName          string                            `json:"AppDisplayName"`
	ConditionalAccessStatus string                            `json:"ConditionalAccessStatus"`
	CorrelationID           string                            `json:"CorrelationId"`
	CreatedDateTime         string                            `json:"CreatedDateTime"`
	ID                      string                            `json:"ID"`
	IPAddress               string                            `json:"IPAddress"`
	IsInteractive           string                            `json:"IsInteractive"`
	ResourceDisplayName     string                            `json:"ResourceDisplayName"`
	ResourceID              string                            `json:"ResourceId"`
	RiskDetail              string                            `json:"RiskDetail"`
	RiskLevelDuringSignIn   string                            `json:"RiskLevelDuringSignIn"`
	RiskState               string                            `json:"RiskState"`
	Source                  string                            `json:"Source"`
}

type DiscoveryUserSoftwareDeviceDetail struct {
	Browser         string `json:"Browser"`
	IsCompliant     string `json:"IsCompliant"`
	IsManaged       string `json:"IsManaged"`
	OperatingSystem string `json:"OperatingSystem"`
}

type DiscoveryUserSoftwareLocation struct {
	City            string                                      `json:"City"`
	CountryOrRegion string                                      `json:"CountryOrRegion"`
	State           string                                      `json:"State"`
	GeoCoordinates  DiscoveryUserSoftwareLocationGeoCoordinates `json:"GeoCoordinates"`
}

type DiscoveryUserSoftwareLocationGeoCoordinates struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type DiscoveryUserSoftwareResult struct {
	ErrorCode     string `json:"ErrorCode"`
	FailureReason string `json:"FailureReason"`
	Status        string `json:"Status"`
}

type DiscoveryUserLocation struct {
	Address1 string `json:"Address1"`
	City     string `json:"City"`
	Country  string `json:"Country"`
	Office   string `json:"Office"`
	ZipCode  string `json:"ZipCode"`
}

type DiscoveryUserLocationGeoCoordinates struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type DiscoveryUserSoftwareAssignedApp struct {
	ApplicationID          string `json:"ApplicationID"`
	AlternateID            string `json:"AlternateID"`
	ApplicationDisplayName string `json:"ApplicationDisplayName"`
	PrincipalType          string `json:"PrincipalType"`
	Source                 string `json:"Source"`
}

type DiscoveryUserSoftwareApplicationTags struct {
	ApplicationID string `json:"ApplicationID"`
	Tag           string `json:"Tag"`
}