package api_models

type TenantRequest struct {
	TenantId string `json:"tenantId"`
}

type TenantResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
