package errors

// Error levels
const (
	Debug   = 5
	Info    = 4
	Warning = 3
	Error   = 2
	Fatal   = 1
)

// Error codes
var (
	ErrorDatabase = createErrorDetails(1, "Something went wrong, please contact an administrator", Fatal)
	ErrorRabbitMQ = createErrorDetails(2, "Something went wrong, please contact an administrator [RMQ]", Fatal)
	ErrorElastic  = createErrorDetails(3, "Something went wrong, please contact an administrator [ES]", Fatal)
	ErrorGeneric  = createErrorDetails(4, "%s", Warning)
)
