package api

// ReqLogin ...
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
