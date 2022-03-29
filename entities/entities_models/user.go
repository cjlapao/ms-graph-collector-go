package entities_models

type User struct {
	ID                   string `json:"id" bson:"_id"`
	AccountEnabled       bool   `json:"accountEnabled" bson:"accountEnabled"`
	CreatedDateTime      string `json:"createdDatetime" bson:"createdFatetime"`
	Email                string `json:"email" bson:"email"`
	JobTitle             string `json:"jobTitle" bson:"jobTitle"`
	Surname              string `json:"surname" bson:"surname"`
	GivenName            string `json:"givenName" bson:"givenName"`
	UserPrincipalName    string `json:"userPrincipalName" bson:"userPrincipalName"`
	PreferredLanguage    string `json:"preferredLanguage" bson:"preferredLanguage"`
	DisplayName          string `json:"displayName" bson:"displayName"`
	AssignedLicenses     int64  `json:"assignedLicenses" bson:"assignedLicenses"`
	AssignedServicePlans int64  `json:"assignedServicePlans" bson:"assignedServicePlans"`
}
