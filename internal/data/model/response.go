package model

type RegProfileResponseOK struct {
	// A response returned from the server.
	// e.g. "Successfully registered"
	Message string `json:"message"`
}

type RegProfileResponseFailure struct {
	// A response returned from the server.
	// e.g. "No records were found"
	Message string `json:"message"`
}
