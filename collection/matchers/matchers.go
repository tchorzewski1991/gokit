package matchers

// Bool is a predefines Matcher for bool values.
func Bool(target bool) func(elem bool) bool {
	return func(elem bool) bool {
		return elem == target
	}
}

// Str is a predefines Matcher for string values.
func Str(target string) func(elem string) bool {
	return func(elem string) bool {
		return elem == target
	}
}

// Int is a predefines Matcher for integer like values.
func Int[T ~int | ~int8 | ~int16 | ~int32 | ~int64](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}

// Float is a predefines Matcher for float like values.
func Float[T ~float32 | ~float64](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}
