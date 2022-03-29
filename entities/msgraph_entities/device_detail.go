package msgraph_entities

type DeviceDetail struct {
	DeviceID        string `json:"deviceId"`
	DisplayName     string `json:"displayName"`
	OperatingSystem string `json:"operatingSystem"`
	Browser         string `json:"browser"`
	IsCompliant     bool   `json:"isCompliant"`
	IsManaged       bool   `json:"isManaged"`
	TrustType       string `json:"trustType"`
}
