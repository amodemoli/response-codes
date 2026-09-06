package response

// Code is a specialized type to prevent the use of invalid codes.
// Use this type for functions that require a response code.
type Code string

// Success codes (2xx)
const (
	Ok       Code = "OK"        // Request succeeded
	Created  Code = "CREATED"   // Resource created
	Accepted Code = "ACCEPTED"  // Request accepted but not yet processed
	NoContent Code = "NO_CONTENT" // Successful request with no response body
	Healthly Code = "HEALTHLY"  // Service is healthy (for health checks)
)

// Client error codes (4xx)
const (
	BadRequest          Code = "BAD_REQUEST"           // Invalid request syntax or parameters
	Unauthorized        Code = "UNAUTHORIZED"          // Missing or invalid authentication
	Forbidden           Code = "FORBIDDEN"             // Authenticated but not authorized
	NotFound            Code = "NOT_FOUND"             // Requested resource does not exist
	MethodNotAllowed    Code = "METHOD_NOT_ALLOWED"    // HTTP method not supported
	Conflict            Code = "CONFLICT"              // Request conflicts with current state (e.g., duplicate)
	UnprocessableEntity Code = "UNPROCESSABLE_ENTITY"  // Request valid but cannot be processed (validation)
	TooManyRequests     Code = "TOO_MANY_REQUESTS"     // Rate limit exceeded
)

// Server error codes (5xx)
const (
	InternalServerError Code = "INTERNAL_SERVER_ERROR" // Generic server error
	NotImplemented      Code = "NOT_IMPLEMENTED"       // Feature not implemented
	BadGateway          Code = "BAD_GATEWAY"           // Invalid response from upstream
	ServiceUnavailable  Code = "SERVICE_UNAVAILABLE"   // Service is temporarily down
	GatewayTimeout      Code = "GATEWAY_TIMEOUT"       // Upstream request timed out
)

// Business / domain specific errors
const (
	// Authentication & authorization
	UserNotFound       Code = "USER_NOT_FOUND"       // User not found
	InvalidCredentials Code = "INVALID_CREDENTIALS"  // Wrong username/password
	InvalidToken       Code = "INVALID_TOKEN"        // Token is invalid
	TokenExpired       Code = "TOKEN_EXPIRED"        // Token has expired
	PermissionDenied   Code = "PERMISSION_DENIED"    // Insufficient permissions

	// Validation & input
	InvalidInput      Code = "INVALID_INPUT"        // General invalid input
	ValidationFailed  Code = "VALIDATION_FAILED"    // Validation failure
	CannotValidate    Code = "CANNOT_VALIDATE"      // Cannot perform validation

	// Duplicate & conflict
	EmailAlreadyExists    Code = "EMAIL_ALREADY_EXISTS"    // Email already registered
	UsernameAlreadyExists Code = "USERNAME_ALREADY_EXISTS" // Username already taken
	DuplicateEntry        Code = "DUPLICATE_ENTRY"         // Duplicate entry
	AlreadyExists         Code = "ALREADY_EXISTS"          // Resource already exists

	// Resource & operation
	ResourceNotFound   Code = "RESOURCE_NOT_FOUND"   // Requested resource does not exist
	ActionNotAllowed   Code = "ACTION_NOT_ALLOWED"   // Action not permitted
	OperationFailed    Code = "OPERATION_FAILED"     // Operation failed
	DependencyFailed   Code = "DEPENDENCY_FAILED"    // Dependency error

	// Timeout & cancellation
	Timeout  Code = "TIMEOUT"   // Operation timed out
	Canceled Code = "CANCELED"  // Operation canceled

	// Generic / fallback
	InvalidArgument Code = "INVALID_ARGUMENT" // Invalid argument provided
	OutOfRange      Code = "OUT_OF_RANGE"     // Value out of allowed range
	Unauthenticated Code = "UNAUTHENTICATED"  // User not authenticated
	Aborted         Code = "ABORTED"          // Operation aborted
	DataLoss        Code = "DATA_LOSS"        // Data loss occurred
	Unknown         Code = "UNKNOWN"          // Unknown error
)

// HTTPStatusMapping maps each response code to its corresponding HTTP status.
var HTTPStatusMapping = map[Code]int{
	Ok:                  200,
	Created:             201,
	Accepted:            202,
	NoContent:           204,
	Healthly:            200,
	BadRequest:          400,
	Unauthorized:        401,
	Forbidden:           403,
	NotFound:            404,
	MethodNotAllowed:    405,
	Conflict:            409,
	UnprocessableEntity: 422,
	TooManyRequests:     429,
	InternalServerError: 500,
	NotImplemented:      501,
	BadGateway:          502,
	ServiceUnavailable:  503,
	GatewayTimeout:      504,
	UserNotFound:        404,
	InvalidCredentials:  401,
	InvalidToken:        401,
	TokenExpired:        401,
	PermissionDenied:    403,
	InvalidInput:        400,
	ValidationFailed:    422,
	CannotValidate:      422,
	EmailAlreadyExists:  409,
	UsernameAlreadyExists: 409,
	DuplicateEntry:      409,
	AlreadyExists:       409,
	ResourceNotFound:    404,
	ActionNotAllowed:    403,
	OperationFailed:     500,
	DependencyFailed:    503,
	Timeout:             504,
	Canceled:            499,
	InvalidArgument:     400,
	OutOfRange:          400,
	Unauthenticated:     401,
	Aborted:             409,
	DataLoss:            500,
	Unknown:             500,
}

// ToHTTPStatus returns the HTTP status code for a given response code.
// If the code is not found, it returns 500 (Internal Server Error).
func (c Code) ToHTTPStatus() int {
	if status, ok := HTTPStatusMapping[c]; ok {
		return status
	}
	return 500
}

// 🐋 This file was generated with the assistance of DeepSeek (DeepSeek AI model)
// to save development time and accelerate delivery. It is intended to provide
// a comprehensive set of response codes suitable for both simple websites and
// microservices, ensuring clarity and consistency across all responses.
