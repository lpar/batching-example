package main

// The object-oriented solution

type Batcher[V any] struct {
	Size   int
	Index  int
	Values []V
}

// NewBatcher creates a new Batcher with the specified size.
func NewBatcher[V any](values []V, size int) *Batcher[V] {
	return &Batcher[V]{
		Index:  0,
		Values: values,
		Size:   size,
	}
}

// NextBatch returns the next batch of values from the Batcher.
func (b *Batcher[V]) NextBatch() ([]V, bool) {
	if b.Values == nil || b.Size <= 0 {
		return nil, false
	}
	if b.Index >= len(b.Values) {
		return nil, false
	}
	end := b.Index + b.Size
	if end > len(b.Values) {
		end = len(b.Values)
	}
	batch := b.Values[b.Index:end]
	b.Index = end
	return batch, true
}
