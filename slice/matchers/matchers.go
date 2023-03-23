package matchers

import "golang.org/x/exp/constraints"

// Str is a predefined slice.MatchFunc for string like values.
func Str[T ~string](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}

// Int is a predefined slice.MatchFunc for integer like values.
func Int[T constraints.Integer](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}

// Float is a predefined slice.MatchFunc for float like values.
func Float[T constraints.Float](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}

// Bool is a predefined slice.MatchFunc for bool like values.
func Bool[T ~bool](target T) func(elem T) bool {
	return func(elem T) bool {
		return elem == target
	}
}
