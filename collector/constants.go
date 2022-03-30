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
	TopRecords = 999
)

const (
	MSGraphMicrosoftApplicationsCollection = "Microsoft.Graph.MicrosoftApplications"
	MSGraphUsersApplicationsCollection     = "Microsoft.Graph.UsersApplications"
	MSGraphUsersCollection                 = "Microsoft.Graph.Users"
	MSGraphUsageCollection                 = "Microsoft.Graph.Signins"
	MSGraphServicePrincipalsCollection     = "Microsoft.Graph.ServicePrincipals"
	NeuronsUsers                           = "Neurons.Users"
)
