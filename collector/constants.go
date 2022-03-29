package collector

const (
	LoginBaseUrl                       = "https://login.microsoftonline.com"
	GraphBaseUrl                       = "https://graph.microsoft.com"
	TokenUrl                           = "/oauth2/v2.0/token"
	ApplicationsUrl                    = "/beta/applications"
	UsersUrl                           = "/beta/users"
	SigninsUrl                         = "/beta/auditlogs/signins"
	SubscribedSkusUrl                  = "/beta/subscribedSkus"
	ServicePrincipalUrl                = "/beta/servicePrincipals"
	Office365ActivationsUserCountsUrl  = "/beta/reports/getOffice365ActivationsUserCounts?$format=application/json"
	Office365ActivationCountsUrl       = "/beta/reports/getOffice365ActivationCounts?$format=application/json"
	Office365ActivationsUserDetailsUrl = "/beta/reports/getOffice365ActivationsUserDetail?$format=application/json"
	DirectoryObjectUrl                 = "/beta/directoryObjects"
)

const (
	ODataTop = "?$top=999"
)

const (
	MSGraphOrgApplicationsCollection                = "Microsoft.Graph.OrgApplications"
	MSGraphUsersCollection                          = "Microsoft.Graph.Users"
	MSGraphUsersOrgApplicationsCollection           = "Microsoft.Graph.Users.OrgApplications"
	MSGraphUsersLicensesCollection                  = "Microsoft.Graph.Users.Licenses"
	MSGraphUsageCollection                          = "Microsoft.Graph.Usages"
	MSGraphServicePrincipalsCollection              = "Microsoft.Graph.ServicePrincipals"
	MSGraphTenantSkusCollection                     = "Microsoft.Graph.Tenant.Skus"
	MSGraphOffice365ActivationsUserCountsCollection = "Microsoft.Graph.Office365.ActivationsUserCounts"
	MSGraphOffice365ActivationsUserDetailCollection = "Microsoft.Graph.Office365.ActivationsUserDetails"
	MSGraphOffice365ActivationCountsCollection      = "Microsoft.Graph.Office365.ActivationCounts"
	MSGraphOffice365AppUserDetailCollection         = "Microsoft.Graph.Office365.Office365AppUserDetail"
)
