package entities_models

type UserApplication struct {
	ID                            string `json:"id" bson:"_id"`
	UserId                        string `json:"userId" bson:"userId"`
	UserServicePrincipal          string `json:"userServicePrincipal" bson:"userServicePrincipal"`
	UserDisplayName               string `json:"userDisplayName" bson:"userDisplayName"`
	ApplicationId                 string `json:"applicationId" bson:"applicationId"`
	ApplicationServicePrincipalId string `json:"applicationServicePrincipalId" bson:"applicationServicePrincipalId"`
	ApplicationDisplayName        string `json:"applicationDisplayName" bson:"applicationDisplayName"`
	Type                          string `json:"type" bson:"type"`
	Recognized                    bool   `json:"recognized" bson:"recognized"`
	RecognizedSource              string `json:"recognizedSource" bson:"recognizedSource"`
	LastUsed                      string `json:"lastUsed" bson:"lastUsed"`
	LastClientAppUsed             string `json:"lastClientAppUsed" bson:"lastClientAppUsed"`
	LastOperatingSystem           string `json:"lastOperatingSystem" bson:"lastOperatingSystem"`
	LastBrowser                   string `json:"lastBrowser" bson:"lastBrowser"`
	LastIpAddress                 string `json:"lastIpAddress" bson:"lastIpAddress"`
	LastIsInteractive             string `json:"lastIsInteractive" bson:"lastIsInteractive"`
	LastCity                      string `json:"lastCity" bson:"lastCity"`
	LastState                     string `json:"lastState" bson:"lastState"`
	LastCountryOrRegion           string `json:"lastCountryOrRegion" bson:"lastCountryOrRegion"`
	LoginCount                    int64  `json:"loginCount" bson:"loginCount"`
	ProductId                     string `json:"productId" bson:"productId"`
	FriendlyName                  string `json:"friendlyName" bson:"friendlyName"`
}

type UserApplicationWithDetails struct {
	ID                            string      `json:"id" bson:"_id"`
	UserId                        string      `json:"userId" bson:"userId"`
	UserServicePrincipal          string      `json:"userServicePrincipal" bson:"userServicePrincipal"`
	UserDisplayName               string      `json:"userDisplayName" bson:"userDisplayName"`
	ApplicationId                 string      `json:"applicationId" bson:"applicationId"`
	ApplicationServicePrincipalId string      `json:"applicationServicePrincipalId" bson:"applicationServicePrincipalId"`
	ApplicationDisplayName        string      `json:"applicationDisplayName" bson:"applicationDisplayName"`
	Type                          string      `json:"type" bson:"type"`
	Recognized                    bool        `json:"recognized" bson:"recognized"`
	RecognizedSource              string      `json:"recognizedSource" bson:"recognizedSource"`
	LastUsed                      string      `json:"lastUsed" bson:"lastUsed"`
	LastClientAppUsed             string      `json:"lastClientAppUsed" bson:"lastClientAppUsed"`
	LastOperatingSystem           string      `json:"lastOperatingSystem" bson:"lastOperatingSystem"`
	LastBrowser                   string      `json:"lastBrowser" bson:"lastBrowser"`
	LastIpAddress                 string      `json:"lastIpAddress" bson:"lastIpAddress"`
	LastIsInteractive             string      `json:"lastIsInteractive" bson:"lastIsInteractive"`
	LastCity                      string      `json:"lastCity" bson:"lastCity"`
	LastState                     string      `json:"lastState" bson:"lastState"`
	LastCountryOrRegion           string      `json:"lastCountryOrRegion" bson:"lastCountryOrRegion"`
	LoginCount                    int64       `json:"loginCount" bson:"loginCount"`
	User                          User        `json:"user" bson:"user"`
	Application                   Application `json:"application" bson:"application"`
}
