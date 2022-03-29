package msgraph_entities

type Location struct {
	City            string         `json:"city"`
	State           string         `json:"state"`
	CountryOrRegion string         `json:"countryOrRegion"`
	GeoCoordinates  GeoCoordinates `json:"geoCoordinates"`
}
