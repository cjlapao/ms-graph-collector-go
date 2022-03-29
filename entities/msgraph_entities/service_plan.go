package msgraph_entities

type ServicePlan struct {
	ServicePlanID      string             `json:"servicePlanId" bson:"servicePlanId"`
	ServicePlanName    string             `json:"servicePlanName" bson:"servicePlanName"`
	ProvisioningStatus ProvisioningStatus `json:"provisioningStatus" bson:"provisioningStatus"`
	AppliesTo          AppliesTo          `json:"appliesTo" bson:"appliesTo"`
}
