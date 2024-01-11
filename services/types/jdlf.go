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
