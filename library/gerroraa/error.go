package gerroraa

// Error is custom error for additional features.
type Error struct {
	error error  // Wrapped error.
	stack stack  // Stack array, which records the stack information when this error is created or wrapped.
	text  string // Error text, which is created by New* functions.
	code  int    // Error code if necessary.
}