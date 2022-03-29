package entities_models

type ActiveApplication struct {
	ID          string `json:"id" bson:"_id"`
	DisplayName string `json:"displayName" bson:"displayName"`
	Last30Days  int32  `json:"last30Days" bson:"last30Days"`
	Last90Days  int32  `json:"last90Days" bson:"last90Days"`
	Users       []ActiveApplicationUser
}

type ActiveApplicationUser struct {
	ID             string `json:"id" bson:"_id"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	LastActiveDate string `json:"lastActiveDate" bson:"lastActiveDate"`
}
