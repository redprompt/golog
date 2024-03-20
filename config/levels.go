package config

// Define available levels
type Level int

// Logging levels
const (
	Fatal Level = iota
	Silent
	Error
	Info
	Warning
	Verbose
	Debug
)

// Returns iota levels as string
func (l Level) String() string {
	return [...]string{"fatal", "silent", "error", "info", "warning", "verbose", "debug"}[l]
}
