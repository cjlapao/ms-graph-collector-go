package summaries_entities

type UsersSummary struct {
	Count            int64 `json:"count"`
	ActiveLast30Days int64 `json:"activeLast30Days"`
	ActiveLast90Days int64 `json:"activeLast90Days"`
	Inactive         int64 `json:"inactive"`
}
