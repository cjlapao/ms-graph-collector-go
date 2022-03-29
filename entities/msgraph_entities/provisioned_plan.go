package msgraph_entities

type ProvisionedPlan struct {
	CapabilityStatus   CapabilityStatus   `json:"capabilityStatus"`
	ProvisioningStatus ProvisioningStatus `json:"provisioningStatus"`
	Service            string             `json:"service"`
}
