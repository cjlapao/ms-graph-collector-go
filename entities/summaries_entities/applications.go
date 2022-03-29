package summaries_entities

type ApplicationSummary struct {
	Count                    int64 `json:"count"`
	RecognizedApplications   int64 `json:"recognizedApplications"`
	UnrecognizedApplications int64 `json:"unrecognizedApplications"`
	ActiveApplications       int64 `json:"activeApplications"`
}
