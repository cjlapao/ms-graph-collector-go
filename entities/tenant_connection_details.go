package entities

type TenantConnectionDetails struct {
	Name                 string `json:"name" bson:"name"`
	TenantId             string `json:"tenantId" bson:"tenantId"`
	ClientId             string `json:"clientId" bson:"clientId"`
	ClientSecret         string `json:"clientSecret" bson:"clientSecret"`
	UnoBaseUrl           string `json:"unoBaseUrl" bson:"unoBaseUrl"`
	UnoLoginUrl          string `json:"unoLoginUrl" bson:"unoLoginUrl"`
	NeuronsTenantId      string `json:"neuronsTenantId" bson:"neuronsTenantId"`
	NeuronsTenantUrl     string `json:"neuronsTenantUrl" bson:"neuronsTenantUrl"`
	LoginAppClientSecret string `json:"loginAppClientSecret" bson:"loginAppClientSecret"`
	LoginAppClientId     string `json:"loginAppClientId" bson:"loginAppClientId"`
}
