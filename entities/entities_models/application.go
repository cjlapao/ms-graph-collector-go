package entities_models

type Application struct {
	ID                    string            `json:"id" bson:"_id"`
	DisplayName           string            `json:"displayName" bson:"displayName"`
	DeletedDateTime       string            `json:"deletedDateTime" bson:"deletedDateTime"`
	Enabled               bool              `json:"accountEnabled" bson:"accountEnabled"`
	ApplicationId         string            `json:"appId" bson:"appId"`
	Description           interface{}       `json:"description" bson:"description"`
	ServicePrincipalNames []string          `json:"servicePrincipalNames" bson:"servicePrincipalNames"`
	ServicePrincipalType  string            `json:"servicePrincipalType" bson:"servicePrincipalType"`
	SignInAudience        string            `json:"signInAudience" bson:"signInAudience"`
	Tags                  []string          `json:"tags" bson:"tags"`
	Type                  string            `json:"type" bson:"type"`
	Recognized            bool              `json:"recognized" bson:"recognized"`
	RecognizedSource      string            `json:"recognizedSource" bson:"recognizedSource"`
	Users                 []ApplicationUser `json:"users" bson:"users"`
	LowUsageUsages        []ApplicationUser `json:"lowUsageUsers" bson:"lowUsageUsers"`
	LastActiveDate        string            `json:"lastActiveDate" bson:"lastActiveDate"`
	Usage                 string            `json:"usage" bson:"usage"`
	TotalLoginCount       int64             `json:"totalLoginCount" bson:"totalLoginCount"`
	ProductId             string            `json:"productId" bson:"productId"`
	FriendlyName          string            `json:"friendlyName" bson:"friendlyName"`
}

type ApplicationUser struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
