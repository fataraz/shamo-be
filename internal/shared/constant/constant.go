package constant

import "net/http"

const (
	// FormatDateString ...
	FormatDateString = "2006-01-02 15:04:05"
)

var (
	ErrorGeneralMsg           = "Internal Server Error"
	ErrorBadRequestMsg        = "Bad Request"
	ErrorValidationMsg        = "Validation Error"
	ErrorInvalidRequestMsg    = "Invalid request"
	ErrorNotFoundMsg          = "Data not found"
	ErrorInvalidEmailMsg      = "Invalid email"
	ErrorNotAuthorizedMsg     = "Not authorized"
	ErrorInRequestValidParams = "Invalid Request Params"
)

var (
	ErrorGeneral            = NewError(http.StatusInternalServerError, ErrorGeneralMsg)
	ErrorValidation         = NewError(http.StatusBadRequest, ErrorValidationMsg)
	ErrorInvalidRequest     = NewError(http.StatusBadRequest, ErrorInvalidRequestMsg)
	ErrorDataNotFound       = NewError(http.StatusNotFound, ErrorNotFoundMsg)
	ErrorInvalidEmail       = NewError(http.StatusBadRequest, ErrorInvalidEmailMsg)
	ErrorUserCannotAccess   = NewError(http.StatusUnauthorized, ErrorNotAuthorizedMsg)
	ErrorDatabase           = NewError(http.StatusInternalServerError, "Database Error")
	ErrorUserNotFound       = NewError(http.StatusNotFound, "User Not Found")
	ErrorPasswordNotMatch   = NewError(http.StatusUnauthorized, "Password not match")
	ErrorProductNotFound    = NewError(http.StatusNotFound, "Product Not Found")
	ErrorInvalidPhoneNumber = NewError(http.StatusBadRequest, "Invalid Phone Number")
)

// NewError ...
func NewError(errorCode int, message string) error {
	return &ApplicationError{
		Code:    errorCode,
		Message: message,
	}
}

// ApplicationError ...
type ApplicationError struct {
	Code    int
	Message string
}

// Error ...
func (e *ApplicationError) Error() string {
	return e.Message
}
