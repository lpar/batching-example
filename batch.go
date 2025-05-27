package main

// The just-a-simple-function solution

// Batch segments a slice of values into multiple slices with the specified maximum size.
func Batch[Type any](values []Type, size int) [][]Type {
	if len(values) > 0 && size > 0 {
		batches := make([][]Type, 0, (len(values)+size-1)/size)
		for size < len(values) {
			values, batches = values[size:], append(batches, values[0:size:size])
		}
		return append(batches, values)
	}
	return make([][]Type, 0)
}
