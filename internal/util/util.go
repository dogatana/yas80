package util

// Filter
func Filter[T any](s []T, f func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Count
func Count[T any](s []T, f func(T) bool) int {
	out := 0
	for _, v := range s {
		if f(v) {
			out++
		}
	}
	return out
}
