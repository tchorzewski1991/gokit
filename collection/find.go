package collection

// Find returns the first element of the collection for which Match returns true.
func Find[T comparable](elems []T, match Match[T]) (res T, ok bool) {
	for idx := range elems {
		if match(elems[idx]) {
			return elems[idx], true
		}
	}
	return
}
