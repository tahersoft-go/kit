package errs

import (
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg ...string) ErrorResponse {
	if len(msg) == 0 {
		msg = append(msg, "خطایی رخ داده است. لطفا به تیم پشتیبانی اطلاع دهید")
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg[0],
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg ...string) ErrorResponse {
	if len(msg) == 0 {
		msg = append(msg, "درخواست مورد نظر یافت نشد")
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: msg[0],
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg ...string) ErrorResponse {
	if len(msg) == 0 {
		msg = append(msg, "You are not authenticated to perform the requested action.")
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: msg[0],
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg ...string) ErrorResponse {
	if len(msg) == 0 {
		msg = append(msg, "You are not authorized to perform the requested action.")
	}
	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: msg[0],
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg ...string) ErrorResponse {
	if len(msg) == 0 {
		msg = append(msg, "Your request is in a bad format.")
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg[0],
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InvalidInput creates a new error response representing a data validation error (HTTP 400).
func InvalidInput(errs validation.Errors) ErrorResponse {
	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}

	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "There is some problem with the data you submitted.",
		Details: details,
	}
}
