package summaries_entities

type ActiveApplicationsSummary struct {
	Count      int64 `json:"count"`
	Last30Days int64 `json:"last30Days"`
	Last90Days int64 `json:"last90Days"`
}
