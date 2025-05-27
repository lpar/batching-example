package main

import (
	"slices"
	"testing"
)

var testCases = []struct {
	Name      string
	Input     []int
	BatchSize int
	Output    [][]int
	Calls     int
}{
	{
		Name:      "empty input",
		Input:     []int{},
		BatchSize: 3,
		Output:    [][]int{},
		Calls:     0,
	},
	{
		Name:      "batch size 1",
		Input:     []int{1, 2, 3},
		BatchSize: 1,
		Output:    [][]int{{1}, {2}, {3}},
		Calls:     3,
	},
	{
		Name:      "batch size 2",
		Input:     []int{1, 2, 3, 4, 5},
		BatchSize: 2,
		Output:    [][]int{{1, 2}, {3, 4}, {5}},
		Calls:     3,
	},
	{
		Name:      "batch size 0",
		Input:     []int{1, 2, 3, 4, 5},
		BatchSize: 0,
		Output:    [][]int{},
		Calls:     0,
	},
	{
		Name:      "batch size bigger than data",
		Input:     []int{1, 2, 3, 4, 5},
		BatchSize: 10,
		Output:    [][]int{{1, 2, 3, 4, 5}},
		Calls:     1,
	},
}

func TestBatch(t *testing.T) {
	checkResults := func(t *testing.T, got [][]int, expected [][]int) {
		t.Helper()
		if len(got) != len(expected) {
			t.Fatalf("got %d results, expected %d", len(got), len(expected))
		}
		for i := range got {
			if len(got[i]) != len(expected[i]) {
				t.Errorf("got %d, expected %d", len(got[i]), len(expected[i]))
			}
			for j := range got[i] {
				if got[i][j] != expected[i][j] {
					t.Errorf("got %d, expected %d", got[i][j], expected[i][j])
				}
			}
		}
	}

	for _, tc := range testCases {

		// Test regular batch
		t.Run("regular "+tc.Name, func(t *testing.T) {
			got := Batch(tc.Input, tc.BatchSize)
			checkResults(t, got, tc.Output)
		})

		// Test functional batch
		t.Run("functional "+tc.Name, func(t *testing.T) {
			got := [][]int{}
			calls := 0
			f := func(batch []int) error {
				got = append(got, batch)
				calls++
				return nil
			}
			if err := BatchFunc(tc.Input, tc.BatchSize, f); err != nil {
				t.Fatalf("error: %v", err)
			}
			if calls != tc.Calls {
				t.Fatalf("called %d times, expected %d", calls, tc.Calls)
			}
			checkResults(t, got, tc.Output)
		})

		// Test OO batch
		t.Run("oo "+tc.Name, func(t *testing.T) {
			got := [][]int{}
			batcher := NewBatcher(tc.Input, tc.BatchSize)
			for {
				batch, ok := batcher.NextBatch()
				if !ok {
					break
				}
				got = append(got, batch)
			}
			checkResults(t, got, tc.Output)
		})

		// Test iterator batch
		t.Run("iterator "+tc.Name, func(t *testing.T) {
			got := [][]int{}
			calls := 0
			f := func(batch []int) bool {
				got = append(got, batch)
				calls++
				return true
			}
			BatchSeq(tc.Input, tc.BatchSize)(f)
			if calls != tc.Calls {
				t.Fatalf("called %d times, expected %d", calls, tc.Calls)
			}
			checkResults(t, got, tc.Output)
		})

		// Test slices.Chunk
		t.Run("slices.Chunk "+tc.Name, func(t *testing.T) {
			if tc.BatchSize <= 0 {
				t.Skip("slices.Chunk does not support batch size 0 or negative")
			}
			got := [][]int{}
			calls := 0
			f := func(batch []int) bool {
				got = append(got, batch)
				calls++
				return true
			}
			slices.Chunk(tc.Input, tc.BatchSize)(f)
			if calls != tc.Calls {
				t.Fatalf("called %d times, expected %d", calls, tc.Calls)
			}
			checkResults(t, got, tc.Output)
		})
	}
}

func Test_IterationExample(t *testing.T) {
	// Example usage of the BatchSeq function
	values := []int{1, 2, 3, 4, 5}
	batchSize := 2
	batches := BatchSeq(values, batchSize)
	for batch := range batches {
		t.Log(batch)
	}
}

func BenchmarkBatch(b *testing.B) {
	for b.Loop() {
		Batch(testCases[0].Input, testCases[0].BatchSize)
	}
}

func BenchmarkBatchFunc(b *testing.B) {
	for b.Loop() {
		BatchFunc(testCases[0].Input, testCases[0].BatchSize, func(batch []int) error {
			return nil
		})
	}
}

func BenchmarkBatcher(b *testing.B) {
	for b.Loop() {
		batcher := NewBatcher(testCases[0].Input, testCases[0].BatchSize)
		for {
			if _, ok := batcher.NextBatch(); !ok {
				break
			}
		}
	}
}

func BenchmarkBatchSeq(b *testing.B) {
	for b.Loop() {
		BatchSeq(testCases[0].Input, testCases[0].BatchSize)(func(batch []int) bool {
			return true
		})
	}
}

func BenchmarkChunk(b *testing.B) {
	for b.Loop() {
		slices.Chunk(testCases[0].Input, testCases[0].BatchSize)(func(batch []int) bool {
			return true
		})
	}
}
