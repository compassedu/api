package types

type AuthenticationResult struct {
	D struct {
		TwoFAuthRequired bool                    `json:"2FAuthRequired"`
		FriendlyMessage  string                  `json:"friendlyMessage"`
		Success          bool                    `json:"success"`
		TechnicalMessage string                  `json:"technicalMessage"`
		Roles            []AuthenticatedUserRole `json:"roles"`
	} `json:"d"`
}
type AuthenticatedUserRole struct {
	BaseRole float64 `json:"baseRole"`
	Fqdn     string  `json:"fqdn"`
	SchoolId string  `json:"schoolId"`
	UserId   float64 `json:"userId"`
}
type CalendarTransport struct {
	ActivityId               float64                 `json:"activity"`
	ActivityImportIdentifier *string                 `json:"activityImportIdentifier"`
	ActivityType             float64                 `json:"activityType"`
	AllDay                   bool                    `json:"allDay"`
	AttendanceMode           float64                 `json:"attendanceMode"`
	AttendeeUserId           float64                 `json:"attendeeUserId"`
	BackgroundColor          string                  `json:"backgroundColor"`
	CalendarId               *float64                `json:"calendarId"`
	CategoryIds              *[]float64              `json:"categoryIds"`
	Comment                  *string                 `json:"comment"`
	Description              string                  `json:"description"`
	EventSetupStatus         *float64                `json:"eventSetupStatus"`
	Finish                   string                  `json:"finish"`
	Guid                     string                  `json:"guid"`
	InClassStatus            *float64                `json:"inClassStatus"`
	InstanceId               string                  `json:"instanceId"`
	IsRecurring              bool                    `json:"isRecurring"`
	LearningTaskId           *float64                `json:"learningTaskId"`
	LessonPlanConfigured     bool                    `json:"lessonPlanConfigured"`
	Location                 *float64                `json:"location"`
	Locations                []CalendarEventLocation `json:"locations"`
	LongTitle                string                  `json:"longTitle"`
	LongTitleWithoutTime     string                  `json:"longTitleWithoutTime"`
	ManagerId                float64                 `json:"managerId"`
	Managers                 []CalendarEventManager  `json:"managers"`
	MinutesMeetingId         *float64                `json:"minutesMeetingId"`
	Period                   string                  `json:"period"`
	RecurringFinish          *string                 `json:"recurringFinish"`
	RecurringStart           *string                 `json:"recurringStart"`
	RepeatDays               *int                    `json:"repeatDays"`
	RepeatForever            bool                    `json:"repeatForever"`
	RepeatFrequency          float64                 `json:"repeatFrequency"`
	RepeatUntil              *string                 `json:"repeatUntil"`
	RollMarked               bool                    `json:"rollMarked"`
	RunningStatus            float64                 `json:"runningStatus"`
	Start                    string                  `json:"start"`
	TargetStudentId          float64                 `json:"targetStudentId"`
	TeachingDaysOnly         bool                    `json:"teachingDaysOnly"`
	TextColor                string                  `json:"textColor"`
	Title                    string                  `json:"title"`
	UnavailablePd            *string                 `json:"unavailablePd"`
}
type CalendarEventLocation struct {
	CoveringLocationId   *float64 `json:"coveringLocationId"`
	CoveringLocationName *string  `json:"coveringLocationName"`
	LocationId           float64  `json:"locationId"`
	LocationName         string   `json:"locationName"`
}
type CalendarEventManager struct {
	CoveringImportIdentifier *string  `json:"coveringImportIdentifier"`
	CoveringUserId           *float64 `json:"coveringUserId"`
	ManagerImportIdentifier  string   `json:"managerImportIdentifier"`
	ManagerUserId            float64  `json:"managerUserId"`
}
type User struct {
	BaseRole                float64    `json:"baseRole"`
	CampusId                float64    `json:"campusId"`
	Ce                      string     `json:"ce"`
	DisplayCode             string     `json:"displayCode"`
	DoNotContact            bool       `json:"doNotContact"`
	F                       string     `json:"f"`
	Finish                  *string    `json:"finish"`
	FirstName               string     `json:"fn"`
	GovtCode1               string     `json:"govtCode1"`
	GovtCode2               string     `json:"govtCode2"`
	H                       *string    `json:"h"`
	HasRegisteredDevice     bool       `json:"hasRegisteredDevice"`
	Id                      float64    `json:"id"`
	ImportIdentifier        string     `json:"ii"`
	LastName                string     `json:"ln"`
	MobileNumber            string     `json:"mobileNumber"`
	Name                    string     `json:"n"`
	NameFirstPrefLastIdForm string     `json:"nameFirstPrefLastIdForm"`
	NamePrefFirst           string     `json:"namePrefFirst"`
	NamePrefLastId          string     `json:"namePrefLastId"`
	FirstNameImportId       string     `json:"nif"`
	Ns                      string     `json:"ns"`
	Picture                 string     `json:"p"`
	ParentIds               *[]float64 `json:"parentIds"`
	PhoneExtension          string     `json:"phoneExtension"`
	ProfilePictureV         string     `json:"pv"`
	ReportName              string     `json:"reportName"`
	Rfid                    string     `json:"rfid"`
	RfidCards               []float64  `json:"rfidCards"`
	Start                   string     `json:"start"`
	SussiId                 string     `json:"sussiId"`
	U                       string     `json:"u"`
	UserStatus              float64    `json:"userStatus"`
}
type LC struct {
	Archived bool    `json:"archived"`
	Building string  `json:"building"`
	Id       float64 `json:"id"`
	LongName string  `json:"longName"`
	Name     string  `json:"n"`
	RoomName string  `json:"roomName"`
}
type AlertItem struct {
	AlertItemId    float64 `json:"AlertItemId"`
	Body           string  `json:"Body"`
	Content        string  `json:"Content"`
	Dismissible    bool    `json:"Dismissible"`
	ImageSourceUrl string  `json:"ImageSourceUrl"`
	IsWarning      bool    `json:"IsWarning"`
	LinkText       string  `json:"LinkText"`
	LinkUrl        string  `json:"LinkUrl"`
	Title          string  `json:"Title"`
	Type           float64 `json:"Type"`
}
type StandardClass struct {
	AttendanceModeDefault   float64  `json:"attendanceModeDefault"`
	CampusId                *float64 `json:"campusId"`
	CheckInEnabledDefault   float64  `json:"checkInEnabledDefault"`
	CustomLocation          *string  `json:"customLocation"`
	Description             *string  `json:"description"`
	ExtendedStatusId        float64  `json:"extendedStatusId"`
	ExternalIntegrationCode string   `json:"externalIntegrationCode"`
	FacultyName             string   `json:"facultyName"`
	Finish                  string   `json:"finish"`
	HaparaSyncEnabled       bool     `json:"haparaSyncEnabled"`
	Id                      float64  `json:"id"`
	ImportIdentifier        string   `json:"importIdentifier"`
	ImportTeachers          bool     `json:"importTeachers"`
	LayerAllowsImport       bool     `json:"layerAllowsImport"`
	LayerId                 float64  `json:"layerId"`
	LocationId              *float64 `json:"locationId"`
	ManagerId               float64  `json:"managerId"`
	ManagerImportIdentifier string   `json:"managerImportIdentifier"`
	Name                    string   `json:"name"`
	RollTapThreshold        float64  `json:"rollTapThreshold"`
	Start                   string   `json:"start"`
	SubjectId               float64  `json:"subjectId"`
	SubjectImportIdentifier string   `json:"subjectImportIdentifier"`
	SubjectLongName         string   `json:"subjectLongName"`
	TimetableStructureId    float64  `json:"timetableStructureId"`
	YearLevelShortName      string   `json:"yearLevelShortName"`
}
type NewsItem struct {
	CommunicationType      float64              `json:"CommunicationType"`
	Content1               string               `json:"Content1"`
	Content2               *string              `json:"Content2"`
	CreatedByAdmin         bool                 `json:"CreatedByAdmin"`
	EmailSentDate          string               `json:"EmailSentDate"`
	Finish                 string               `json:"Finish"`
	Locked                 bool                 `json:"Locked"`
	NewsItemId             string               `json:"NewsItemId"`
	NotificationStatus     float64              `json:"NotificationStatus"`
	PostDateTime           string               `json:"PostDateTime"`
	Priority               bool                 `json:"Priority"`
	PublicWebsite          bool                 `json:"PublicWebsite"`
	PublishToLinkedSchools bool                 `json:"PublishToLinkedSchools"`
	PublishToTalkingPoints bool                 `json:"PublishToTalkingPoints"`
	SenderId               float64              `json:"SenderId"`
	ShowImagesFullScreen   bool                 `json:"ShowImagesFullScreen"`
	Start                  string               `json:"Start"`
	StartFinishString      string               `json:"StartFinishString"`
	Title                  string               `json:"Title"`
	UserId                 float64              `json:"UserId"`
	UserImageUrl           string               `json:"UserImageUrl"`
	UserName               string               `json:"UserName"`
	Attachments            []NewsitemAttachment `json:"Attachments"`
}
type NewsitemAttachment struct {
	AssetId              float64  `json:"AssetId"`
	FileAssetType        float64  `json:"FileAssetType"`
	IsImage              bool     `json:"IsImage"`
	Name                 string   `json:"Name"`
	OriginalFileName     string   `json:"OriginalFileName"`
	SourceOrganisationId *float64 `json:"SourceOrganisationId"`
	UiLink               string   `json:"UiLink"`
	Url                  *string  `json:"Url"`
}

// type ActionCentreEvent struct {
// 	additionalContactDetails
// ""
// additionalDetails
// ""
// administrationDetails
// ""
// allowConsentWithoutPayment
// false
// allowDecline
// false
// amountPaid
// 0
// attendeeLimit
// null
// attendeeStatus
// 1
// canProvideEventConsent
// true
// confirmedAttendeesCount
// 51
// consentDt
// "2024-01-10T23 15 51Z"
// consentFormId
// 2
// consentName
// "DARAGH O'GORMAN"
// consentPaymentDue
// "2024-02-01T23 59 59Z"
// consentReturnLocation
// "na"
// cost
// 0
// description
// null
// dressCode
// ""
// educativePurpose
// "Detailed itinerary will be provided closer to the trip."
// eligibleForCSEF
// false
// expenses
// null
// faculty
// ""
// fimsUnallocatedFunds
// 0
// finish
// "2024-02-13T07 55 00Z"
// id
// 5912
// isOptIn
// false
// isParentAttendee
// false
// isPartiallyPaid
// false
// location
// null
// medicalDetails
// {__type  "ActionCentreEventContactDetails http //jdlf.com.au/ns/data/actioncentre/web",…}
// additionalMedicalDetails
// ""
// __type
// "ActionCentreEventContactDetails http //jdlf.com.au/ns/data/actioncentre/web"
// name
// "Ski Trip"
// onlineProcessing
// true
// paidConsentMethod
// 1
// paidViaCsef
// false
// paidViaPaymentPlan
// false
// parentAttendeeId
// 0
// parentAttendeeName
// null
// paymentDt
// null
// paymentPlanInstalments
// null
// questions
// []
// refundCutoffDt
// null
// refundStatus
// null
// requiresConsent
// true
// requiresPayment
// false
// sessions
// [{__type  "ActionCentreEventSession http //jdlf.com.au/ns/data/actioncentre/web", campusName  "",…}]
// showActionPlans
// false
// start
// "2024-02-08T03 30 00Z"
// studentConsentContent
// "I give permission for Killian O'Gorman to attend this event. Where the staff member in charge is unable to contact me, or where it is impracticable to contact me, I authorise the staff member in charge to 1) consent to any medical or surgical attention deemed necessary by a medical practitioner, and 2) administer such first-aid as the staff member in charge judges to be reasonably necessary. I understand that this is an official school event and that Killian O'Gorman will adhere to the dress code and behave in alignment with the school&#39;s code of behaviour and the house rules.&nbsp;"
// studentContactInfo
// []
// studentEmergencyContacts
// []
// studentId
// 4225
// studentMedicalInfo
// []
// studentName
// "O'Gorman, Killian"
// transport
// ""
// useCsef
// false
// usePaymentPlans
// false
// __type
// "ActionCentreEvent http //jdlf.com.au/ns/data/actioncentre/web"
// }
