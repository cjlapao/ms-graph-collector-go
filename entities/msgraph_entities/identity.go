package msgraph_entities

type Identity struct {
	SignInType       string `json:"signInType"`
	Issuer           string `json:"issuer"`
	IssuerAssignedID string `json:"issuerAssignedId"`
}
