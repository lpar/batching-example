
Different ways of implementing batching of slice values.

- `batch.go`: Straightforward procedural utility function with loop, returning slice of slices.
- `batcher.go`: Object oriented batcher with a method to pull the next batch.
- `batchfunc.go`: Functional approach, pass a function that gets called with each batch.
- `batchiter.go`: Using Go's `iter` package.
- `batch_test.go`: Unit tests and benchmarks.
