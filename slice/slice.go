package slice

// MatchFunc provides unified signature used to verify whether given condition has been met.
// Dig into collection/matchers to take advantage of predefined matching funcs.
type MatchFunc[T any] func(elem T) bool

// Find returns the first element of the collection for which MatchFunc returns true.
func Find[T any](elems []T, fn MatchFunc[T]) (res T, ok bool) {
	for _, elem := range elems {
		if fn(elem) {
			return elem, true
		}
	}
	return
}
