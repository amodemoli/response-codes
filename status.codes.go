package status

// Status is a specialized type to prevent the use of invalid status values.
// Use this type for functions that require a response status.
type Status string

// Common status values for API responses
const (
	Ok      Status = "Ok"      // Operation completed successfully
	Error   Status = "Error"   // Operation failed with an error
	Ready   Status = "Ready"   // Service or resource is ready
	Pending Status = "Pending" // Operation is in progress (not yet completed)
	Failed  Status = "Failed"  // Operation failed permanently
	Warning Status = "Warning" // Operation completed with warnings
	Info    Status = "Info"    // Informational status (no action needed)
	Unknown Status = "Unknown" // Status cannot be determined
)

// Additional statuses for more granular control
const (
	Created   Status = "Created"   // Resource was created successfully
	Updated   Status = "Updated"   // Resource was updated successfully
	Deleted   Status = "Deleted"   // Resource was deleted successfully
	Accepted  Status = "Accepted"  // Request accepted but processing is deferred
	Canceled  Status = "Canceled"  // Operation was canceled
	Timeout   Status = "Timeout"   // Operation timed out
	Aborted   Status = "Aborted"   // Operation was aborted
	Invalid   Status = "Invalid"   // Request or data is invalid
	NotFound  Status = "NotFound"  // Requested resource was not found
	Duplicate Status = "Duplicate" // Duplicate entry or conflict
	Forbidden Status = "Forbidden" // Access forbidden
	Unauthorized Status = "Unauthorized" // Authentication required or failed
)

// 🐋 This file was generated with the assistance of DeepSeek (DeepSeek AI model)
// to save development time and accelerate delivery. It provides a clear and
// concise set of status indicators suitable for both simple websites and
// microservices, ensuring consistent communication of the overall response state.
