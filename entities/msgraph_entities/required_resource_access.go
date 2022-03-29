package msgraph_entities

type RequiredResourceAccess struct {
	ResourceAppID  string           `json:"resourceAppId"`
	ResourceAccess []ResourceAccess `json:"resourceAccess"`
}
