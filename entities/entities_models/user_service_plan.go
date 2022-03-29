package entities_models

import "github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"

type UserServicePlan struct {
	ID                 string                              `json:"id" bson:"_id"`
	ServicePlanID      string                              `json:"servicePlanId" bson:"servicePlanId"`
	UserID             string                              `json:"userId" bson:"userId"`
	Service            string                              `json:"service" bson:"service"`
	ProvisioningStatus msgraph_entities.ProvisioningStatus `json:"provisioningStatus" bson:"provisioningStatus"`
	Capabilitystatus   msgraph_entities.ProvisioningStatus `json:"capabilityStatus" bson:"capabilityStatus"`
	AppliesTo          msgraph_entities.AppliesTo          `json:"appliesTo" bson:"appliesTo"`
	AssignedToUser     bool                                `json:"assignedToUser" bson:"assignedToUser"`
	ProductId          string                              `json:"productId" bson:"productId"`
	Recognized         bool                                `json:"recognized" bson:"recognized"`
	FriendlyName       string                              `json:"friendlyName" bson:"friendlyName"`
}
