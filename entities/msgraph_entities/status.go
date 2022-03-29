package msgraph_entities

type Status struct {
	ErrorCode         int64       `json:"errorCode"`
	FailureReason     string      `json:"failureReason"`
	AdditionalDetails interface{} `json:"additionalDetails"`
}
