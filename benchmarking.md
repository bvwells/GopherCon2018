# Benchmarking

The document details how to run Go benchmarking.

Create a benchmark in a ```_test.go``` file. For example:

```
func BenchmarkSum(b *testing.B) {
	...
```

Run the benchmark by executing the following command:

```
go build -gcflags -S benchmarking_test.go
```

# Dump assembly

```
go build -gcflags -S main.go
```
