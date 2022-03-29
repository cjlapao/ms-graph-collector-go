package entities_models

type CollectorUserAssignedApplication struct {
	AlternateID                    string `json:"AlternateID" bson:"alternateId"`
	ApplicationDisplayName         string `json:"ApplicationDisplayName" bson:"applicationDisplayName"`
	ApplicationID                  string `json:"ApplicationID" bson:"applicationId"`
	ApplicationOwnerOrganizationID string `json:"ApplicationOwnerOrganizationID" bson:"applicationOwnerOrganizationID"`
	PrincipalType                  string `json:"PrincipalType" bson:"principalType"`
	Source                         string `json:"Source" bson:"source"`
}
