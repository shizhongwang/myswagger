package handlers

import "fmt"

// AppError is a generic error message returned by a server
// Generic error message returned as a string
// swagger:response errorResponse
type AppError struct {
	StatusCode 	int `json:"status_code"`
	Message 	string `json:"message,omitempty"`
	ID 			string `json:"id,omitempty"`
	RequestID 	string `json:"request_id,omitempty"`
}

// Generic error message returned as a string
// swagger:response error errorForbidden
type errorForbidden struct {
	// Description of the error
	// in: body
	Body AppError
}

// ErrInvalidChatroomPath is an error message when the ChatroomRequest path is not valid
var ErrInvalidChatroomPath = fmt.Errorf("Invalid Path, path should be /Chatrooms/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}



// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}
