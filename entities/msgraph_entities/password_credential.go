package msgraph_entities

type PasswordCredential struct {
	CustomKeyIdentifier interface{} `json:"customKeyIdentifier"`
	DisplayName         string      `json:"displayName"`
	EndDateTime         string      `json:"endDateTime"`
	Hint                string      `json:"hint"`
	KeyID               string      `json:"keyId"`
	SecretText          interface{} `json:"secretText"`
	StartDateTime       string      `json:"startDateTime"`
}
