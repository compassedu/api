package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/compassedu/api/services/login"
	"github.com/compassedu/api/services/types"
)

// Login authenticates a user with Compass Education API and returns session cookies and user ID.
//
// The function takes the provided 'username', 'password', and 'schoolId' to authenticate the user.
// It utilizes the Compass Education API to perform the authentication.
//
// Parameters:
//   - username: The username of the user.
//   - password: The password of the user.
//   - schoolId: The unique identifier for the school.
//
// Returns:
//   - cookies: A semicolon-separated string containing session cookies.
//   - userId:   The user ID obtained after successful authentication.
//   - err:      An error, if any, that occurred during the authentication process.
//
// Example:
//
//	cookies, userId, err := Login("exampleUsername", "examplePassword", "exampleSchoolID")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use cookies and userId as needed.
func loginFunc(username string, password string, schoolId string) (cookies string, userId float64, err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", 0, err
	}
	client := &http.Client{Jar: jar}
	credentials := login.LoginCredentials{
		SessionState: "readonly",
		Username:     username,
		Password:     password,
	}
	reqbody, err := json.Marshal(credentials)
	if err != nil {
		return "", 0, err
	}
	url := "https://" + schoolId + ".compass.education/services/admin.svc/AuthenticateUserCredentials"
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqbody))
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	// if res.StatusCode != 200 {
	// 	return "", 0, errors.New("HTTP Response code doesn't equal 200")
	// }
	var u types.AuthenticationResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return "", 0, err
	}
	if u.D.Success == false {
		return "", 0, errors.New("Incorrect Credentials")
	}
	if u.D.TwoFAuthRequired == true {
		return "", 0, errors.New("Two Factor Auth Not Supported Yet")
	}
	c := res.Cookies()
	var cs []string
	for _, cookie := range c {
		cs = append(cs, cookie.Name+"="+cookie.Value)
	}
	return strings.Join(cs, ";"), u.D.Roles[0].UserId, nil
}

// Login authenticates a user with Compass Education API and returns session cookies and user ID.
//
// The function takes the provided 'username', 'password', and 'schoolId' to authenticate the user.
// It utilizes the Compass Education API to perform the authentication.
//
// Parameters:
//   - username: The username of the user.
//   - password: The password of the user.
//   - schoolId: The unique identifier for the school.
//
// Returns:
//   - cookies: A semicolon-separated string containing session cookies.
//   - userId:   The user ID obtained after successful authentication.
//   - err:      An error, if any, that occurred during the authentication process.
//
// Example:
//
//	cookies, userId, err := Login("exampleUsername", "examplePassword", "exampleSchoolID")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use cookies and userId as needed.
func Login(username string, password string, schoolId string) (cookies string, userId float64, err error) {
	// Runing Function Once
	loginFunc(username, password, schoolId)
	// Second Run
	c, u, e := loginFunc(username, password, schoolId)
	return c, u, e
}
