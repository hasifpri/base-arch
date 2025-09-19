package helperexception

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"reflect"
	"regexp"
	"unicode"
)

type ErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Code string

const (
	InvalidArgumentCode  Code = "INVALID_ARGUMENT"  // Represents an invalid argument error.
	NotFoundCode         Code = "NOT_FOUND"         // Represents a not found error.
	AlreadyExistsCode    Code = "ALREADY_EXISTS"    // Represents an already exists error.
	PermissionDeniedCode Code = "PERMISSION_DENIED" // Represents a permission denied error.
	UnauthenticatedCode  Code = "UNAUTHENTICATED"   // Represents an unauthenticated error.
	InternalErrorCode    Code = "INTERNAL"          // Represents an internal error.
)

type Exception struct {
	Code    Code
	Message any
	Error   error
}

func (e *Exception) GetError() *string {
	if e.Error != nil {
		err := e.Error.Error()
		return &err
	}
	return nil
}

func (e *Exception) GetHttpCode() int {
	switch e.Code {
	case InvalidArgumentCode:
		return 400
	case NotFoundCode:
		return 404
	case AlreadyExistsCode:
		return 409
	case PermissionDeniedCode:
		return 403
	case UnauthenticatedCode:
		return 401
	case InternalErrorCode:
		return 500
	default:
		return 500
	}
}

func (e *Exception) IsEqual(err *Exception) bool {
	if err == nil {
		return e == nil
	}
	if e == nil {
		return err == nil
	}
	return e.Code == err.Code && reflect.DeepEqual(e.Message, err.Message)
}

func InvalidArgument(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    InvalidArgumentCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func NotFound(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    NotFoundCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func AlreadyExists(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func PermissionDenied(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    PermissionDeniedCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func Unauthenticated(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    UnauthenticatedCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func Internal(message any, err error) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    InternalErrorCode,
		Message: errorMessage,
		Error:   err,
	}
}

func Conflict(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func convertToString(message any) string {
	if msg, ok := message.(string); ok {
		return msg
	}
	return fmt.Sprintf("%v", message)
}

func TranslateMessage(err error) interface{} {
	var errorsResponse []ErrorResponse
	var message, fieldName string

	if err.Error() == "mongo: no documents in result" { // mongo
		return "No data found"
	} else if err.Error() == "record not found" { // gorm
		return "No data found"
	}

	// Handle PostgreSQL error
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			// Unique constraint violation
			column, value := extractKeyValueFromDetail(pgErr.Detail)
			fieldName = snakeCase(column)
			message = fmt.Sprintf("Duplicate key error: %s already exists", value)
		case "23503":
			// Foreign key violation
			fieldName = snakeCase(pgErr.ColumnName)
			message = fmt.Sprintf("Foreign key constraint error: related record not found or linked to %s", pgErr.ConstraintName)
		case "23502":
			// Not null violation
			fieldName = snakeCase(pgErr.ColumnName)
			message = fmt.Sprintf("Not null constraint error: field %s cannot be null", pgErr.ColumnName)
		case "23514":
			// Check constraint violation
			fieldName = snakeCase(pgErr.ColumnName)
			message = fmt.Sprintf("Check constraint error: %s", pgErr.ConstraintName)
		case "23504":
			// Exclusion constraint violation
			fieldName = snakeCase(pgErr.ColumnName)
			message = fmt.Sprintf("Exclusion constraint error: conflict with existing data due to %s", pgErr.ConstraintName)
		default:
			// Other PostgreSQL errors can be handled generally
			message = fmt.Sprintf("Database error: %s", pgErr.Message)
		}
		errorsResponse = append(errorsResponse, ErrorResponse{
			Field:   fieldName,
			Message: message,
		})

	} else if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			fieldName = snakeCase(e.Field())
			tag := e.Tag()

			switch tag {
			case "required":
				message = fmt.Sprintf("%s is required", fieldName)
			default:
				message = fmt.Sprintf("%s failed validation on the '%s' tag", fieldName, tag)
			}

			errorsResponse = append(errorsResponse, ErrorResponse{
				Field:   fieldName,
				Message: message,
			})
		}
	} else {
		errorsResponse = append(errorsResponse, ErrorResponse{
			Message: err.Error(),
		})
	}
	return errorsResponse

}

func snakeCase(input string) string {
	var result []rune
	for i, r := range input {
		if unicode.IsUpper(r) {
			// Add an underscore before the uppercase letter if it's not the first character
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func extractKeyValueFromDetail(detail string) (column, value string) {
	re := regexp.MustCompile(`Key \(([^)]+)\)=\(([^)]+)\)`)
	matches := re.FindStringSubmatch(detail)
	if len(matches) == 3 {
		return matches[1], matches[2]
	}
	return "", ""
}
