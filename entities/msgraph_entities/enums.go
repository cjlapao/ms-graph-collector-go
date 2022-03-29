package msgraph_entities

type ProvisioningStatus string

const (
	ProvisioningStatusPendingActivation   ProvisioningStatus = "PendingActivation"
	ProvisioningStatusPendingInput        ProvisioningStatus = "PendingInput"
	ProvisioningStatusPendingProvisioning ProvisioningStatus = "PendingProvisioning"
	ProvisioningStatusSuccess             ProvisioningStatus = "Success"
	ProvisioningStatusDisabled            ProvisioningStatus = "Disabled"
	ProvisioningStatusEnabled             ProvisioningStatus = "Enabled"
)

type AppliesTo string

const (
	AppliesToCompany AppliesTo = "Company"
	AppliesToUser    AppliesTo = "User"
)
