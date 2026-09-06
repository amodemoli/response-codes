package statuscodes

// A special type to prevent the use of invalid codes, 
// for functions that require a code, name the type this way
type Status string

// Add your custom status codes
const (
	Ok    Status = "Ok"
	Error Status = "Error"
	Ready Status = "Ready"
)
