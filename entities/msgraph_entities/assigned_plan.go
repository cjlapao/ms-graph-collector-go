package msgraph_entities

type AssignedPlan struct {
	AssignedDateTime string             `json:"assignedDateTime"`
	CapabilityStatus ProvisioningStatus `json:"capabilityStatus"`
	Service          string             `json:"service"`
	ServicePlanID    string             `json:"servicePlanId"`
}
