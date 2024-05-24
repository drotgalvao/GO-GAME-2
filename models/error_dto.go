package models

// ErrorDTO represents an error response structure used across the API.
// It includes a code to identify the type of error and a message describing the error.
//
// Example:
//   {
//     "code": 400,
//     "message": "Invalid request parameters."
//   }
type ErrorDTO struct {
    Code    int    `json:"code"` // The error code
    Message string `json:"message"` // The error message
}