package login

type LoginCredentials struct {
	SessionState string `json:"sessionstate"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}
