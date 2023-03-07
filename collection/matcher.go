package collection

// Match provides unified signature used to verify whether given condition has been met.
// Dig into collection/matchers to take advantage of predefined matchers.
type Match[T comparable] func(elem T) bool
