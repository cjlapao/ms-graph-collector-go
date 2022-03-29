package entities_models

type AppRecognitionEntry struct {
	ID                 string   `json:"id" bson:"_id"`
	AppName            string   `json:"appName" bson:"appName"`
	RecognitionPhrases []string `json:"recognitionPhrases" bson:"recognitionPhrases"`
}

type AppRecognitionResult struct {
	ID         string `json:"id" bson:"_id"`
	AppName    string `json:"appName" bson:"appName"`
	Recognized bool   `json:"recognized" bson:"recognized"`
}

type AppRecognitionRequest struct {
	Phrase string `json:"phrase"`
}

type MicrosoftLicenseRecognitionEntry struct {
	ID             string   `json:"id" bson:"_id"`
	Name           string   `json:"name" bson:"name"`
	RecognitionIDs []string `json:"recognitionIds" bson:"RecognitionIds"`
}

type MicrosoftLicenseRecognitionResult struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	Recognized bool   `json:"recognized" bson:"recognized"`
}

type MicrosoftLicenseRecognitionRequest struct {
	Id string `json:"id"`
}
