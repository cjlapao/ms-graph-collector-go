package msgraph_entities

type AppRoleAssignment struct {
	ID                   string `json:"id" bson:"_id"`
	CreationTimestamp    string `json:"creationTimestamp" bson:"creationtimestamp"`
	AppRoleID            string `json:"appRoleId" bson:"approleid"`
	PrincipalDisplayName string `json:"principalDisplayName" bson:"principaldisplayname"`
	PrincipalID          string `json:"principalId" bson:"principalid"`
	PrincipalType        string `json:"principalType" bson:"principaltype"`
	ResourceDisplayName  string `json:"resourceDisplayName" bson:"resourcedisplayname"`
	ResourceID           string `json:"resourceId" bson:"resourceid"`
}
