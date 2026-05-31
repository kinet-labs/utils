package utils

// Zero returns the zero value of any type T.
func Zero[T any]() T {
	var zero T
	return zero
}
