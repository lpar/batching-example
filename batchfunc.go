package main

// The functional solution

// BatchFunc segments a slice of values into multiple slices with the specified maximum size
// and applies a function to each segment. If the function returns an error, the iteration stops
// and the error is returned.
func BatchFunc[T any](ss []T, size int, f func([]T) error) error {
	if len(ss) == 0 || size <= 0 {
		return nil
	}
	start := 0
	for start < len(ss) {
		end := start + size
		if end > len(ss) {
			end = len(ss)
		}
		if err := f(ss[start:end]); err != nil {
			return err
		}
		start = end
	}
	return nil
}
