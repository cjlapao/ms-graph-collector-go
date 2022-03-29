package msgraph_entities

type AssignedLicense struct {
	DisabledPlans []string `json:"disabledPlans"`
	SkuID         string   `json:"skuId"`
}
