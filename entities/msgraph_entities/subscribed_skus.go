package msgraph_entities

type SubscribedSku struct {
	CapabilityStatus string        `json:"capabilityStatus" bson:"capabilityStatus"`
	ConsumedUnits    int64         `json:"consumedUnits" bson:"consumedUnits"`
	ID               string        `json:"id" bson:"_id"`
	SkuID            string        `json:"skuId" bson:"skuId"`
	SkuPartNumber    string        `json:"skuPartNumber" bson:"skuPartNumber"`
	AppliesTo        string        `json:"appliesTo" bson:"appliesTo"`
	PrepaidUnits     PrepaidUnits  `json:"prepaidUnits" bson:"prepaidUnits"`
	ServicePlans     []ServicePlan `json:"servicePlans" bson:"servicePlans"`
}
