package msgraph_entities

type UserLicenses struct {
	ID            string        `json:"id" bson:"licenseId"`
	UserID        string        `json:"_" bson:"_id"`
	SkuID         string        `json:"skuId" bson:"skuId"`
	SkuPartNumber string        `json:"skuPartNumber" bson:"skuPartNumber"`
	ServicePlans  []ServicePlan `json:"servicePlans" bson:"servicePlans"`
}
