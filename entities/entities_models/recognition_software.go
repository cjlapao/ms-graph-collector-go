package entities_models

type RecognitionProduct struct {
	ID                 string              `json:"id" bson:"_id"`
	FriendlyName       string              `json:"friendlyName" bson:"friendlyName"`
	Vendor             string              `json:"vendor" bson:"vendor"`
	VendorCode         string              `json:"vendorCode" bson:"vendorCode"`
	VendorID           string              `json:"vendorId" bson:"vendorId"`
	SoftwareType       string              `json:"type" bson:"type"`
	RecognitionPhrases []RecognitionPhrase `json:"recognitionPhrases" bson:"recognitionPhrases"`
	ParentId           string              `json:"parentId,omitempty" bson:"parentId"`
	IsLicensable       bool                `json:"isLicensable" bson:"isLicensable"`
	IsBundle           bool                `json:"isBundle" bson:"isBundle"`
	IsServicePlan      bool                `json:"isServicePlan" bson:"isServicePlan"`
	IsApplication      bool                `json:"isApplication" bson:"isApplication"`
	LicenseType        string              `json:"licenseType" bson:"licenseType"`
	Version            string              `json:"version" bson:"version"`
	IsSaaS             bool                `json:"isSaaS" bson:"isSaaS"`
	BundeledSoftware   []BundledProducts   `json:"bundeledSoftware" bson:"bundeledSoftware"`
}

type RecognitionPhrase struct {
	Phrase     string `json:"phrase" bson:"phrase"`
	IsFilePath bool   `json:"isFilePath" bson:"isFilePath"`
}

type BundledProducts struct {
	FriendlyName  string `json:"friendlyName" bson:"friendlyName"`
	ProductID     string `json:"productId" bson:"productId"`
	VendorCode    string `json:"vendorCode" bson:"vendorCode"`
	VendorID      string `json:"vendorId" bson:"vendorId"`
	IsServicePlan bool   `json:"isServicePlan" bson:"isServicePlan"`
	IsApplication bool   `json:"isApplication" bson:"isApplication"`
	LicenseType   string `json:"licenseType" bson:"licenseType"`
	Version       string `json:"version" bson:"version"`
	SoftwareType  string `json:"type" bson:"type"`
}

type RecognitionProductResponse struct {
	Recognized    bool   `json:"recognized"`
	ID            string `json:"id" bson:"_id"`
	FriendlyName  string `json:"friendlyName" bson:"friendlyName"`
	Vendor        string `json:"vendor" bson:"vendor"`
	VendorCode    string `json:"vendorCode" bson:"vendorCode"`
	VendorID      string `json:"vendorId" bson:"vendorId"`
	SoftwareType  string `json:"type" bson:"type"`
	ParentId      string `json:"parentId,omitempty" bson:"parentId"`
	IsLicensable  bool   `json:"isLicensable" bson:"isLicensable"`
	IsBundle      bool   `json:"isBundle" bson:"isBundle"`
	IsServicePlan bool   `json:"isServicePlan" bson:"isServicePlan"`
	IsApplication bool   `json:"isApplication" bson:"isApplication"`
	LicenseType   string `json:"licenseType" bson:"licenseType"`
	Version       string `json:"version" bson:"version"`
	IsSaaS        bool   `json:"isSaaS" bson:"isSaaS"`
}
