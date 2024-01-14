package compass

import (
	"time"

	"github.com/compassedu/api/services"
	"github.com/compassedu/api/services/types"
)

type API struct {
	// supplied username
	//
	// e.g. JOHNDOE08
	APIUsername string
	// supplied password
	//
	// e.g. *****
	APIPassword string
	// supplied school prefix
	//
	// e.g. ui
	APISchoolPrefix string
	// the base url that will be requested
	//
	// e.g. ui.compass.education
	BaseURL string
	// the cookies which are created
	APICookies string
	// your user id
	APIUserId float64
}

// New creates a new Compass API client.
func New(username string, password string, schoolId string) (*API, error) {
	c, u, err := services.Login(username, password, schoolId)
	if err != nil {
		return &API{}, err
	}
	api := &API{
		BaseURL:         schoolId + ".compass.education",
		APIUsername:     username,
		APIPassword:     password,
		APISchoolPrefix: schoolId,
		APICookies:      c,
		APIUserId:       u,
	}
	return api, nil
}

// Retrieves staff.
//
// The function takes the provided session 'cookies' & 'schoolId'
// to make a request to the Compass Education API and fetch staff.
//
// Returns:
//   - []types.User: A slice of staff in the form of 'types.User'.
//   - error:                      An error, if any, that occurred during the API request or processing of the response.
func (c *API) GetAllStaff() ([]types.User, error) {
	staff, err := services.GetAllStaff(c.APICookies, c.APISchoolPrefix, c.APIUserId)
	if err != nil {
		return []types.User{}, err
	}
	return staff, nil
}

// Retrieves locations.
//
// The function takes the provided session 'cookies' & 'schoolId'
// to make a request to the Compass Education API and fetch locations.
//
// Returns:
//   - []types.LC: A slice of locations in the form of 'types.LC'.
//   - error:                      An error, if any, that occurred during the API request or processing of the response.
func (c *API) GetAllLocations() ([]types.LC, error) {
	locations, err := services.GetAllLocations(c.APICookies, c.APISchoolPrefix)
	if err != nil {
		return []types.LC{}, err
	}
	return locations, nil
}

// Retrieves calendar events for a specified user within a date range.
//
// The function takes the provided session 'cookies', 'schoolId', 'userId', 'startDate', and 'endDate'
// to make a request to the Compass Education API and fetch calendar events for the specified user within the given date range.
//
// Parameters:
//   - startDate:  The start date of the range for which calendar events are requested.
//   - endDate:    The end date of the range for which calendar events are requested.
//
// Returns:
//   - []types.CalendarTransport: A slice of calendar events in the form of 'types.CalendarTransport'.
//   - error:                      An error, if any, that occurred during the API request or processing of the response.
func (c *API) GetClasses(start time.Time, end time.Time) ([]types.CalendarTransport, error) {
	events, err := services.GetCalendarEventsByUser(c.APICookies, c.APISchoolPrefix, c.APIUserId, start, end)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	return events, nil
}

// func (c *API) GetInstance(instanceId string) ([]types.LC, error) {
// 	events, err := services.GetLessonsByInstanceIdQuick(c.APICookies, c.APISchoolPrefix, instanceId)
// 	if err != nil {
// 		return []types.LC{}, err
// 	}
// 	return events, nil
// }
