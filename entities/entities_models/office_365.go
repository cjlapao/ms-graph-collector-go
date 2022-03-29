package entities_models

type Office365ReportCard struct {
	ID                            string `json:"id" bson:"_id"`
	ReportDate                    string `json:"reportDate" bson:"reportDate"`
	TotalProducts                 int64  `json:"totalProducts" bson:"totalProducts"`
	TotalActivations              int64  `json:"totalActivations" bson:"totalActivations"`
	TotalWindowsActivations       int64  `json:"totalWindowsActivations" bson:"totalWindowsActivations"`
	TotalAndroidActivations       int64  `json:"totalAndroidActivations" bson:"totalAndroidActivations"`
	TotalIosActivations           int64  `json:"totalIosActivations" bson:"totalIosActivations"`
	TotalMacActivations           int64  `json:"totalMacActivations" bson:"totalMacActivations"`
	TotalWindowsMobileActivations int64  `json:"totalWindowsMobileActivations" bson:"totalWindowsMobileActivations"`
	TotalActivatedUsers           int64  `json:"totalActivatedUsers" bson:"totalActivatedUsers"`
	TotalAssignedUsers            int64  `json:"totalAssignedUsers" bson:"totalAssignedUsers"`
}

type Office365ProductReport struct {
	ID                       string                        `json:"id" bson:"_id"`
	Office365ProductID       string                        `json:"office365ProductId" bson:"office365ProductId"`
	ProductId                string                        `json:"productId" bson:"productId"`
	Recognized               bool                          `json:"recognized" bson:"recognized"`
	FriendlyName             string                        `json:"friendlyName" bson:"friendlyName"`
	ReportRefreshDate        string                        `json:"reportRefreshDate" bson:"reportRefreshDate"`
	ProductType              string                        `json:"productType" bson:"productType"`
	Assigned                 int64                         `json:"assigned" bson:"assigned"`
	Activated                int64                         `json:"activated" bson:"activated"`
	SharedComputerActivation int64                         `json:"sharedComputerActivation" bson:"sharedComputerActivation"`
	Licenses                 Office365ProductReportDetails `json:"details" bson:"details"`
	Users                    []Office365ProductReportUser  `json:"users" bson:"users"`
}

type Office365ProductReportDetails struct {
	Windows         int64 `json:"windows" bson:"windows"`
	MAC             int64 `json:"mac" bson:"mac"`
	Android         int64 `json:"android" bson:"android"`
	Ios             int64 `json:"ios" bson:"ios"`
	Windows10Mobile int64 `json:"windows10Mobile" bson:"windows10Mobile"`
}

type Office365ProductReportUser struct {
	UserID            string                              `json:"userId" bson:"userId"`
	UserPrincipalName string                              `json:"userPrincipalName" bson:"userPrincipalName"`
	DisplayName       string                              `json:"displayName" bson:"displayName"`
	Details           []Office365ProductReportUserDetails `json:"details" bson:"details"`
}

type Office365ProductReportUserDetails struct {
	ProductType               string      `json:"productType" bson:"productType"`
	LastActivatedDate         interface{} `json:"lastActivatedDate" bson:"lastActivatedDate"`
	Windows                   int64       `json:"windows" bson:"windows"`
	MAC                       int64       `json:"mac" bson:"mac"`
	Windows10Mobile           int64       `json:"windows10Mobile" bson:"windows10Mobile"`
	Ios                       int64       `json:"ios" bson:"ios"`
	Android                   int64       `json:"android" bson:"android"`
	ActivatedOnSharedComputer bool        `json:"activatedOnSharedComputer" bson:"activatedOnSharedComputer"`
}

type Office365ProductReportUserUsage struct {
	ID                 string                                 `json:"id" bson:"_id"`
	UserID             string                                 `json:"userId" bson:"userId"`
	DisplayName        string                                 `json:"displayName" bson:"displayName"`
	ReportRefreshDate  string                                 `json:"reportRefreshDate" bson:"reportRefreshDate"`
	UserPrincipalName  string                                 `json:"userPrincipalName" bson:"userPrincipalName"`
	LastActivationDate string                                 `json:"lastActivationDate" bson:"lastActivationDate"`
	LastActivityDate   string                                 `json:"lastActivityDate" bson:"lastActivityDate"`
	Details            Office365ProductReportUserUsageDetails `json:"details" bson:"details"`
}

type Office365ProductReportUserUsageNew struct {
	ID                 string `json:"id" bson:"_id"`
	UserID             string `json:"userId" bson:"userId"`
	UserName           string `json:"userName" bson:"userName"`
	ReportRefreshDate  string `json:"reportRefreshDate" bson:"reportRefreshDate"`
	UserPrincipalName  string `json:"userPrincipalName" bson:"userPrincipalName"`
	LastActivationDate string `json:"lastActivationDate" bson:"lastActivationDate"`
	LastActivityDate   string `json:"lastActivityDate" bson:"lastActivityDate"`
	ApplicationName    string `json:"applicationName" bson:"applicationName"`
	Platform           string `json:"platform" bson:"platform"`
}

type Office365ProductReportUserUsageDetails struct {
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

type Office365Product struct {
	ID                       string `json:"id" bson:"_id"`
	ReportRefreshDate        string `json:"reportRefreshDate" bson:"reportRefreshDate"`
	ProductType              string `json:"productType" bson:"productType"`
	Assigned                 int64  `json:"assigned" bson:"assigned"`
	Activated                int64  `json:"activated" bson:"activated"`
	SharedComputerActivation int64  `json:"sharedComputerActivation" bson:"sharedComputerActivation"`
	Windows                  int64  `json:"windows" bson:"windows"`
	MAC                      int64  `json:"mac" bson:"mac"`
	Android                  int64  `json:"android" bson:"android"`
	Ios                      int64  `json:"ios" bson:"ios"`
	Windows10Mobile          int64  `json:"windows10Mobile" bson:"windows10Mobile"`
	TotalUsers               int64  `json:"totalUsers" bson:"totalUsers"`
}

type Office365ProductUser struct {
	UserID            string `json:"userId" bson:"userId"`
	UserPrincipalName string `json:"userPrincipalName" bson:"userPrincipalName"`
	DisplayName       string `json:"displayName" bson:"displayName"`
}
