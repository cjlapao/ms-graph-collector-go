package msgraph_entities

type Web struct {
	HomePageURL           interface{}           `json:"homePageUrl"`
	LogoutURL             interface{}           `json:"logoutUrl"`
	RedirectUris          []interface{}         `json:"redirectUris"`
	ImplicitGrantSettings ImplicitGrantSettings `json:"implicitGrantSettings"`
}
