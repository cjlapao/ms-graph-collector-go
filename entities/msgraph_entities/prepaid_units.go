package msgraph_entities

type PrepaidUnits struct {
	Enabled   int64 `json:"enabled" bson:"enabled"`
	Suspended int64 `json:"suspended" bson:"suspended"`
	Warning   int64 `json:"warning" bson:"warning"`
}
