package msgraph_entities

type ParentalControlSettings struct {
	CountriesBlockedForMinors []interface{} `json:"countriesBlockedForMinors,omitempty"`
	LegalAgeGroupRule         string        `json:"legalAgeGroupRule,omitempty"`
}
