package msgraph_entities

type UserApplication struct {
	ID                        string                   `json:"id" bson:"_id"`
	DeletedDateTime           string                   `json:"deletedDateTime"  bson:"deletedDateTime"`
	AppID                     string                   `json:"appId" bson:"appId"`
	ApplicationTemplateID     interface{}              `json:"applicationTemplateId" bson:"applicationTemplateId"`
	DisabledByMicrosoftStatus interface{}              `json:"disabledByMicrosoftStatus" bson:"disabledByMicrosoftStatus"`
	CreatedDateTime           string                   `json:"createdDateTime" bson:"createdDateTime"`
	DisplayName               string                   `json:"displayName" bson:"displayName"`
	Description               interface{}              `json:"description" bson:"description"`
	GroupMembershipClaims     interface{}              `json:"groupMembershipClaims" bson:"groupMembershipClaims"`
	IdentifierUris            []interface{}            `json:"identifierUris" bson:"identifierUris"`
	IsDeviceOnlyAuthSupported interface{}              `json:"isDeviceOnlyAuthSupported" bson:"isDeviceOnlyAuthSupported"`
	IsFallbackPublicClient    interface{}              `json:"isFallbackPublicClient" bson:"isFallbackPublicClient"`
	Notes                     interface{}              `json:"notes" bson:"notes"`
	PublisherDomain           string                   `json:"publisherDomain" bson:"publisherDomain"`
	SignInAudience            string                   `json:"signInAudience" bson:"signInAudience"`
	Tags                      []interface{}            `json:"tags" bson:"tags"`
	TokenEncryptionKeyID      interface{}              `json:"tokenEncryptionKeyId" bson:"tokenEncryptionKeyId"`
	DefaultRedirectURI        interface{}              `json:"defaultRedirectUri" bson:"defaultRedirectUri"`
	Certification             interface{}              `json:"certification" bson:"certification"`
	OptionalClaims            interface{}              `json:"optionalClaims" bson:"optionalClaims"`
	AddIns                    []interface{}            `json:"addIns" bson:"addIns"`
	API                       API                      `json:"api" bson:"api"`
	AppRoles                  []interface{}            `json:"appRoles" bson:"appRoles"`
	Info                      Info                     `json:"info" bson:"info"`
	KeyCredentials            []interface{}            `json:"keyCredentials" bson:"keyCredentials"`
	ParentalControlSettings   ParentalControlSettings  `json:"parentalControlSettings" bson:"parentalControlSettings"`
	PasswordCredentials       []PasswordCredential     `json:"passwordCredentials" bson:"passwordCredentials"`
	PublicClient              PublicClient             `json:"publicClient" bson:"publicClient"`
	RequiredResourceAccess    []RequiredResourceAccess `json:"requiredResourceAccess" bson:"requiredResourceAccess"`
	VerifiedPublisher         VerifiedPublisher        `json:"verifiedPublisher" bson:"verifiedPublisher"`
	Web                       Web                      `json:"web" bson:"web"`
	SPA                       PublicClient             `json:"spa" bson:"spa"`
}
