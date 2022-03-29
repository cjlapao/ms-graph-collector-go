package msgraph_entities

type ServicePrincipal struct {
	ID                                     string            `json:"id"`
	DeletedDateTime                        string            `json:"deletedDateTime"`
	AccountEnabled                         bool              `json:"accountEnabled"`
	AlternativeNames                       []string          `json:"alternativeNames"`
	AppDisplayName                         string            `json:"appDisplayName"`
	AppDescription                         string            `json:"appDescription"`
	AppID                                  string            `json:"appId"`
	ApplicationTemplateID                  string            `json:"applicationTemplateId"`
	AppOwnerOrganizationID                 string            `json:"appOwnerOrganizationId"`
	AppRoleAssignmentRequired              bool              `json:"appRoleAssignmentRequired"`
	CreatedDateTime                        string            `json:"createdDateTime"`
	Description                            string            `json:"description"`
	DisabledByMicrosoftStatus              string            `json:"disabledByMicrosoftStatus"`
	DisplayName                            string            `json:"displayName"`
	Homepage                               string            `json:"homepage"`
	LoginURL                               string            `json:"loginUrl"`
	LogoutURL                              string            `json:"logoutUrl"`
	Notes                                  string            `json:"notes"`
	NotificationEmailAddresses             []string          `json:"notificationEmailAddresses"`
	PreferredSingleSignOnMode              string            `json:"preferredSingleSignOnMode"`
	PreferredTokenSigningKeyThumbprint     string            `json:"preferredTokenSigningKeyThumbprint"`
	ReplyUrls                              []string          `json:"replyUrls"`
	ServicePrincipalNames                  []string          `json:"servicePrincipalNames"`
	ServicePrincipalType                   string            `json:"servicePrincipalType"`
	SignInAudience                         string            `json:"signInAudience"`
	Tags                                   []string          `json:"tags"`
	TokenEncryptionKeyID                   string            `json:"tokenEncryptionKeyId"`
	ResourceSpecificApplicationPermissions []interface{}     `json:"resourceSpecificApplicationPermissions"`
	SamlSingleSignOnSettings               interface{}       `json:"samlSingleSignOnSettings"`
	VerifiedPublisher                      VerifiedPublisher `json:"verifiedPublisher"`
	AddIns                                 []interface{}     `json:"addIns"`
	AppRoles                               []interface{}     `json:"appRoles"`
	Info                                   Info              `json:"info"`
	KeyCredentials                         []interface{}     `json:"keyCredentials"`
	Oauth2PermissionScopes                 []string          `json:"oauth2PermissionScopes"`
	PasswordCredentials                    []interface{}     `json:"passwordCredentials"`
}

type ServicePrincipalWithBson struct {
	ID                                     string            `json:"id" bson:"_id"`
	DeletedDateTime                        string            `json:"deletedDateTime" bson:"deletedDateTime"`
	AccountEnabled                         bool              `json:"accountEnabled" bson:"accountEnabled"`
	AlternativeNames                       []string          `json:"alternativeNames" bson:"alternativeNames"`
	AppDisplayName                         string            `json:"appDisplayName" bson:"appDisplayName"`
	AppDescription                         string            `json:"appDescription" bson:"appDescription"`
	AppID                                  string            `json:"appId" bson:"appId"`
	ApplicationTemplateID                  string            `json:"applicationTemplateId" bson:"applicationTemplateId"`
	AppOwnerOrganizationID                 string            `json:"appOwnerOrganizationId" bson:"appOwnerOrganizationId"`
	AppRoleAssignmentRequired              bool              `json:"appRoleAssignmentRequired" bson:"appRoleAssignmentRequired"`
	CreatedDateTime                        string            `json:"createdDateTime" bson:"createdDateTime"`
	Description                            string            `json:"description" bson:"description"`
	DisabledByMicrosoftStatus              string            `json:"disabledByMicrosoftStatus" bson:"disabledByMicrosoftStatus"`
	DisplayName                            string            `json:"displayName" bson:"displayName"`
	Homepage                               string            `json:"homepage" bson:"homepage"`
	LoginURL                               string            `json:"loginUrl" bson:"loginUrl"`
	LogoutURL                              string            `json:"logoutUrl" bson:"logoutUrl"`
	Notes                                  string            `json:"notes" bson:"notes"`
	NotificationEmailAddresses             []string          `json:"notificationEmailAddresses" bson:"notificationEmailAddresses"`
	PreferredSingleSignOnMode              string            `json:"preferredSingleSignOnMode" bson:"preferredSingleSignOnMode"`
	PreferredTokenSigningKeyThumbprint     string            `json:"preferredTokenSigningKeyThumbprint" bson:"preferredTokenSigningKeyThumbprint"`
	ReplyUrls                              []string          `json:"replyUrls" bson:"replyUrls"`
	ServicePrincipalNames                  []string          `json:"servicePrincipalNames" bson:"servicePrincipalNames"`
	ServicePrincipalType                   string            `json:"servicePrincipalType" bson:"servicePrincipalType"`
	SignInAudience                         string            `json:"signInAudience" bson:"signInAudience"`
	Tags                                   []string          `json:"tags" bson:"tags"`
	TokenEncryptionKeyID                   string            `json:"tokenEncryptionKeyId" bson:"tokenEncryptionKeyId"`
	ResourceSpecificApplicationPermissions []string          `json:"resourceSpecificApplicationPermissions" bson:"resourceSpecificApplicationPermissions"`
	SamlSingleSignOnSettings               string            `json:"samlSingleSignOnSettings" bson:"samlSingleSignOnSettings"`
	VerifiedPublisher                      VerifiedPublisher `json:"verifiedPublisher" bson:"verifiedPublisher"`
	AddIns                                 []string          `json:"addIns" bson:"addIns"`
	AppRoles                               []string          `json:"appRoles" bson:"appRoles"`
	Info                                   Info              `json:"info" bson:"info"`
	KeyCredentials                         []string          `json:"keyCredentials" bson:"keyCredentials"`
	Oauth2PermissionScopes                 []string          `json:"oauth2PermissionScopes" bson:"oauth2PermissionScopes"`
	PasswordCredentials                    []string          `json:"passwordCredentials" bson:"passwordCredentials"`
}
