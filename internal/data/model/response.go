package model

type RegProfileResponseOK struct {
	// A response returned from the server.
	// e.g. "Successfully registered".
	Message string `json:"message"`
}

type RegProfileResponseBR struct {
	// A response returned from the server.
	Message string `json:"message"`
}

type RegProfileResponseFN struct {
	// A response returned from the server.
	// e.g. "ERROR: duplicate key value violates unique constraint".
	Message string `json:"message"`
}

type LogProfileResponseOK struct {
	// Access token to auth API requests.
	AccessToken string `json:"access_token"`
	// Refresh token to recreate access token.
	RefreshToken string `json:"refresh_token"`
}

type LogProfileResponseBR struct {
	// A response returned from the server.
	// e.g. "No records were found".
	Message string `json:"message"`
}

type GetProfileInfoResponseOK struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
}

type GetProfileInfoResponseBR struct {
	// A response returned from the server.
	Message string `json:"message"`
}
