# About

This ia a simple projeto to test Golang generic performance.

## Prerequisite

Go 1.18

## To benchmark

```shell
go test -bench=.
```

## My benchmark

```
goos: linux
goarch: amd64
pkg: genericbenchmark/genericbenchmark
cpu: Intel(R) Pentium(R) CPU G3250 @ 3.20GHz
BenchmarkGenericReduce1-2       49915670                23.80 ns/op
BenchmarkNonGenericReduce1-2    150604094                7.590 ns/op
BenchmarkGenericReduce2-2       50220974                23.52 ns/op
BenchmarkNonGenericReduce2-2    46526138                24.14 ns/op
PASS
ok      genericbenchmark/genericbenchmark       7.291s
```
