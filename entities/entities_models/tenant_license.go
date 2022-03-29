package entities_models

import "github.com/cjlapao/ms-graph-collector-go/entities/msgraph_entities"

type TenantLicenseCard struct {
	TotalLicenses int64 `json:"totalLicenses" bson:"totalLicenses"`
}

type TenantLicenseResponse struct {
	ID                string `json:"id" bson:"_id"`
	SkuID             string `json:"skuId" bson:"skuId"`
	SkuPartNumber     string `json:"skuPartNumber" bson:"skuPartNumber"`
	Status            string `json:"status" bson:"status"`
	Consumed          int64  `json:"consumed" bson:"consumed"`
	Enabled           int64  `json:"enabled" bson:"enabled"`
	Suspended         int64  `json:"suspended" bson:"suspended"`
	Warning           int64  `json:"warning" bson:"warning"`
	TotalCompanyPlans int64  `json:"totalCompanyPlans" bson:"totalCompanyPlans"`
	TotalUserPlans    int64  `json:"totalUserPlans" bson:"totalUserPlans"`
	TotalUsers        int64  `json:"totalUsers" bson:"totalUsers"`
	ProductId         string `json:"productId" bson:"productId"`
	Recognized        bool   `json:"recognized" bson:"recognized"`
	FriendlyName      string `json:"friendlyName" bson:"friendlyName"`
}

type TenantLicense struct {
	ID                  string                     `json:"id" bson:"_id"`
	SkuID               string                     `json:"skuId" bson:"skuId"`
	SkuPartNumber       string                     `json:"skuPartNumber" bson:"skuPartNumber"`
	Status              string                     `json:"status" bson:"status"`
	PrepaidUnits        TenantLicensePrepaidUnit   `json:"prepaidUnits" bson:"prepaidUnits"`
	CompanyServicePlans []TenantLicenseServicePlan `json:"companyServicePlans" bson:"companyServicePlans"`
	UserServicePlans    []TenantLicenseServicePlan `json:"userServicePlans" bson:"userServicePlans"`
	Users               []TenantLicenseUser        `json:"users" bson:"users"`
	ProductId           string                     `json:"productId" bson:"productId"`
	Recognized          bool                       `json:"recognized" bson:"recognized"`
	FriendlyName        string                     `json:"friendlyName" bson:"friendlyName"`
}

type TenantLicensePrepaidUnit struct {
	Consumed  int64 `json:"consumed" bson:"consumed"`
	Enabled   int64 `json:"enabled" bson:"enabled"`
	Suspended int64 `json:"suspended" bson:"suspended"`
	Warning   int64 `json:"warning" bson:"warning"`
}

type TenantLicenseServicePlan struct {
	ServicePlanId      string                              `json:"servicePlanId" bson:"servicePlanId"`
	ServicePlanName    string                              `json:"servicePlanName" bson:"servicePlanName"`
	ProvisioningStatus msgraph_entities.ProvisioningStatus `json:"provisioningStatus" bson:"provisioningStatus"`
	ProductId          string                              `json:"productId" bson:"productId"`
	Recognized         bool                                `json:"recognized" bson:"recognized"`
	FriendlyName       string                              `json:"friendlyName" bson:"friendlyName"`
}

type TenantLicenseUser struct {
	ID                  string `json:"userId" bson:"userId"`
	DisplayName         string `json:"displayName" bson:"displayName"`
	Enabled             int64  `json:"enabled" bson:"enabled"`
	Disabled            int64  `json:"disabled" bson:"disabled"`
	PendingActivation   int64  `json:"pendingActivation" bson:"pendingActivation"`
	PendingInput        int64  `json:"pendingInput" bson:"pendingInput"`
	PendingProvisioning int64  `json:"pendingProvisioning" bson:"pendingProvisioning"`
}
