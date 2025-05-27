package main

import "iter"

// The iter solution

// BatchSeq returns an iterator that produces slices of values of the appropriate size from the input slice.
func BatchSeq[V any](values []V, size int) iter.Seq[[]V] {
	return func(yield func([]V) bool) {
		if len(values) > 0 && size > 0 {
			for start := 0; start < len(values); {
				end := start + size
				if end > len(values) {
					end = len(values)
				}
				if !yield(values[start:end]) {
					break
				}
				start = end
			}
		}
	}
}
