package msgraph_entities

type ImplicitGrantSettings struct {
	EnableAccessTokenIssuance bool `json:"enableAccessTokenIssuance"`
	EnableIDTokenIssuance     bool `json:"enableIdTokenIssuance"`
}
