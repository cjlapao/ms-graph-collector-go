package msgraph_entities

type Office365ActivationsUserCount struct {
	ReportRefreshDate        string `json:"reportRefreshDate" bson:"reportRefreshDate"`
	ProductType              string `json:"productType" bson:"productType"`
	Assigned                 int64  `json:"assigned" bson:"assigned"`
	Activated                int64  `json:"activated" bson:"activated"`
	SharedComputerActivation int64  `json:"sharedComputerActivation" bson:"sharedComputerActivation"`
}

type Office365ActivationCount struct {
	ReportRefreshDate string `json:"reportRefreshDate" bson:"reportRefreshDate"`
	ProductType       string `json:"productType" bson:"productType"`
	Windows           int64  `json:"windows" bson:"windows"`
	MAC               int64  `json:"mac" bson:"mac"`
	Android           int64  `json:"android" bson:"android"`
	Ios               int64  `json:"ios" bson:"ios"`
	Windows10Mobile   int64  `json:"windows10Mobile" bson:"windows10Mobile"`
}

type Office365ActivationsUserDetail struct {
	ReportRefreshDate    string                         `json:"reportRefreshDate" bson:"reportRefreshDate"`
	UserPrincipalName    string                         `json:"userPrincipalName" bson:"userPrincipalName"`
	DisplayName          string                         `json:"displayName" bson:"displayName"`
	UserActivationCounts []Office365UserActivationCount `json:"userActivationCounts" bson:"userActivationCounts"`
}

type Office365UserActivationCount struct {
	ProductType               string      `json:"productType" bson:"productType"`
	LastActivatedDate         interface{} `json:"lastActivatedDate" bson:"lastActivatedDate"`
	Windows                   int64       `json:"windows" bson:"windows"`
	MAC                       int64       `json:"mac" bson:"mac"`
	Windows10Mobile           int64       `json:"windows10Mobile" bson:"windows10Mobile"`
	Ios                       int64       `json:"ios" bson:"ios"`
	Android                   int64       `json:"android" bson:"android"`
	ActivatedOnSharedComputer bool        `json:"activatedOnSharedComputer" bson:"activatedOnSharedComputer"`
}

type Office365AppUserDetail struct {
	ReportRefreshDate  string                    `json:"reportRefreshDate" bson:"reportRefreshDate"`
	UserPrincipalName  string                    `json:"userPrincipalName" bson:"userPrincipalName"`
	LastActivationDate string                    `json:"lastActivationDate" bson:"lastActivationDate"`
	LastActivityDate   string                    `json:"lastActivityDate" bson:"lastActivityDate"`
	Details            []Office365AppUserDetails `json:"details" bson:"details"`
}

type Office365AppUserDetails struct {
	ReportPeriod      int64 `json:"reportPeriod" bson:"reportPeriod"`
	Windows           bool  `json:"windows" bson:"windows"`
	MAC               bool  `json:"mac" bson:"mac"`
	Mobile            bool  `json:"mobile" bson:"mobile"`
	Web               bool  `json:"web" bson:"web"`
	Outlook           bool  `json:"outlook" son:"outlook"`
	Word              bool  `json:"word" bson:"word"`
	Excel             bool  `json:"excel" bson:"excel"`
	PowerPoint        bool  `json:"powerPoint" bson:"powerPoint"`
	OneNote           bool  `json:"oneNote" bson:"oneNote"`
	Teams             bool  `json:"teams"`
	OutlookWindows    bool  `json:"outlookWindows" bson:"outlookWindows"`
	WordWindows       bool  `json:"wordWindows" bson:"wordWindows"`
	ExcelWindows      bool  `json:"excelWindows" bson:"excelWindows"`
	PowerPointWindows bool  `json:"powerPointWindows" bson:"powerPointWindows"`
	OneNoteWindows    bool  `json:"oneNoteWindows" bson:"oneNoteWindows"`
	TeamsWindows      bool  `json:"teamsWindows" bson:"teamsWindows"`
	OutlookMAC        bool  `json:"outlookMac" bson:"outlookMac"`
	WordMAC           bool  `json:"wordMac" bson:"wordMac"`
	ExcelMAC          bool  `json:"excelMac" bson:"excelMac"`
	PowerPointMAC     bool  `json:"powerPointMac" bson:"powerPointMac"`
	OneNoteMAC        bool  `json:"oneNoteMac" bson:"oneNoteMac"`
	TeamsMAC          bool  `json:"teamsMac" bson:"teamsMac"`
	OutlookMobile     bool  `json:"outlookMobile" bson:"outlookMobile"`
	WordMobile        bool  `json:"wordMobile" bson:"wordMobile"`
	ExcelMobile       bool  `json:"excelMobile" bson:"excelMobile"`
	PowerPointMobile  bool  `json:"powerPointMobile" bson:"powerPointMobile"`
	OneNoteMobile     bool  `json:"oneNoteMobile" bson:"oneNoteMobile"`
	TeamsMobile       bool  `json:"teamsMobile" bson:"teamsMobile"`
	OutlookWeb        bool  `json:"outlookWeb" bson:"outlookWeb"`
	WordWeb           bool  `json:"wordWeb" bson:"wordWeb"`
	ExcelWeb          bool  `json:"excelWeb" bson:"excelWeb"`
	PowerPointWeb     bool  `json:"powerPointWeb" bson:"powerPointWeb"`
	OneNoteWeb        bool  `json:"oneNoteWeb" bson:"oneNoteWeb"`
	TeamsWeb          bool  `json:"teamsWeb" bson:"teamsWeb"`
}
