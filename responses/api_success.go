package responses

// APISuccess represents an error that can be sent in an error response.
type APISuccess struct {
	// Status represents the HTTP status code
	Status int `json:"200"`
	// Message is the error message that may be displayed to end users
	Message string `json:"ok"`
	// Data specifies the additional response information
	Data interface{} `json:"data,omitempty"`
}
