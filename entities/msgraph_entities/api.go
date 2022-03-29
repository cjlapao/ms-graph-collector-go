package msgraph_entities

type API struct {
	AcceptMappedClaims          interface{}   `json:"acceptMappedClaims"`
	KnownClientApplications     []interface{} `json:"knownClientApplications"`
	RequestedAccessTokenVersion interface{}   `json:"requestedAccessTokenVersion"`
	Oauth2PermissionScopes      []interface{} `json:"oauth2PermissionScopes"`
	PreAuthorizedApplications   []interface{} `json:"preAuthorizedApplications"`
}
