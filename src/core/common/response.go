package common

import "net/http"

const (
	// GeneralServiceUnavailable General
	GeneralServiceUnavailable = 500

	// GeneralBadRequest General
	GeneralBadRequest   = 400
	GeneralUnauthorized = 401

	ExistedEmailCode    = 400001
	InvalidPasswordCode = 400002
)

const (
	// ExistedEmailMessage message
	ExistedEmailMessage    = "Email is existed"
	InvalidPasswordMessage = "Invalid email/password"
)

// ErrorResponse error response struct
type ErrorResponse struct {
	HTTPCode    int
	ServiceCode int
	Message     string
}

var errorResponseMap = map[int]ErrorResponse{
	GeneralServiceUnavailable: {
		HTTPCode:    http.StatusInternalServerError,
		ServiceCode: GeneralServiceUnavailable,
		Message:     "Service unavailable",
	},
	GeneralBadRequest: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: GeneralBadRequest,
		Message:     "Bad request",
	},
	ExistedEmailCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: ExistedEmailCode,
		Message:     ExistedEmailMessage,
	},
	InvalidPasswordCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: InvalidPasswordCode,
		Message:     InvalidPasswordMessage,
	},
	GeneralUnauthorized: {
		HTTPCode:    http.StatusUnauthorized,
		ServiceCode: GeneralUnauthorized,
		Message:     "Unauthorized",
	},
}

// GetErrorResponse get error response from code
func GetErrorResponse(code int) ErrorResponse {
	if val, ok := errorResponseMap[code]; ok {
		return val
	}

	// default response
	return ErrorResponse{
		HTTPCode:    http.StatusInternalServerError,
		ServiceCode: code,
		Message:     http.StatusText(http.StatusInternalServerError),
	}
}
