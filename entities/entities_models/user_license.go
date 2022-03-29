package entities_models

type UserLicense struct {
	ID            string `json:"id" bson:"_id"`
	LicenseID     string `json:"licenseId" bson:"licenseId"`
	UserID        string `json:"userId" bson:"userId"`
	SkuID         string `json:"skuId" bson:"skuId"`
	SkuPartNumber string `json:"skuPartNumber" bson:"skuPartNumber"`
	DisplayName   string `json:"displayName" bson:"displayName"`
	ProductId     string `json:"productId" bson:"productId"`
	Recognized    bool   `json:"recognized" bson:"recognized"`
	FriendlyName  string `json:"friendlyName" bson:"friendlyName"`
}
