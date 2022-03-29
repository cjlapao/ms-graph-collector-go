package msgraph_entities

type LastUsageResult struct {
	ID      LastUsageId `json:"_id" bson:"_id"`
	Entries []LastUsage `json:"entries" bson:"entries"`
}

type LastUsageId struct {
	AppID  string `json:"appid" bson:"appid"`
	UserID string `json:"userid" bson:"userid"`
}

type LastUsage struct {
	LastUsedDate    string `bson:"lastUsedDate"`
	OperatingSystem string `bson:"operatingSystem"`
	Browser         string `bson:"browser"`
	IpAddress       string `bson:"ipAddress"`
	IsInteractive   string `bson:"isInteractive"`
	City            string `bson:"city"`
	State           string `bson:"state"`
	CountryOrRegion string `bson:"countryOrRegion"`
	ClientAppUsed   string `bson:"clientAppUsed"`
}
