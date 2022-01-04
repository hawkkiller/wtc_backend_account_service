package model

type RegProfileRequest struct {
	// user's email
	Email string `json:"email"`
	// user's password
	Password string `json:"password"`
	// user's username
	Username string `json:"username"`
}

type LogProfileRequest struct {
	// user's email
	Email string `json:"email"`
	// user's password
	Password string `json:"password"`
}
