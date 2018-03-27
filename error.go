package finance

import "encoding/json"

// ErrorCode error code returned by api.
type ErrorCode string

// ErrorDescription detailed error description returned by api.
type ErrorDescription string

const (
	// ErrorCodeArguments is triggered when an
	// exception related to a missing query
	// parameter occurs. Check Description
	// for more information.
	ErrorCodeArguments ErrorCode = "argument-error"

	// ErrorDescriptionSymbols describes a possible
	// error scenario where the symbols parameter was
	// missing.
	ErrorDescriptionSymbols ErrorDescription = "Missing value for the \"symbols\" argument"
)

// ErrorResponse represents an error returned as a response.
type ErrorResponse struct {
	*APIResponse `json:"error"`
}

// Error is the response returned when a call is unsuccessful.
type Error struct {
	Code        ErrorCode        `json:"code"`
	Description ErrorDescription `json:"description"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
