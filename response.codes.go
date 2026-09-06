package responsecodes

// A special type to prevent the use of invalid codes, 
// for functions that require a code, name the type this way
type Code string

// Add your custom response codes
const (
	// success response codes
	Ok       Code = "OK"
	Healthly Code = "HEALTHLY"
	// errors status codes
	ServerErr      Code = "SERVER_ERROR"
	ToManyRequests Code = "TO_MANY_REQUESTS"
	NotFound       Code = "NOT_FOUND"
	BadRequest     Code = "BAD_REQUEST"
)
