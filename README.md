
This code gives examples of different ways of implementing batching of slice values.
It includes the newest option, using Go's `iter` package along with `slices.Chunk` (Go 1.23+).
It then has other examples of using `iter`.

- `batch.go`: Straightforward procedural utility function with loop, returning slice of slices.
- `batcher.go`: Object oriented batcher with a method to pull the next batch.
- `batchfunc.go`: Functional approach, pass a function that gets called with each batch.
- `batchiter.go`: Using Go's `iter` package.
- `batch_test.go`: Unit tests and benchmarks to compare performance of different approaches.
- `iter_examples_test.go`: More examples of how to use Go's `iter` package.
