package responses

//import "encoding/json"

// APISuccess represents an error that can be sent in an error response.
type APISuccess struct {
	// Status represents the HTTP status code
	Status int `json:"status"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message"`
	// Data specifies the additional response information
	Data struct{} `json:"data,omitempty"`
}

// http://choly.ca/post/go-json-marshalling/
//func (s *APISuccess) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&APISuccess{
//		Status:  200,
//		Message: "ok",
//	})
//}