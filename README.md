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
pkg: genericbenchmark
cpu: Intel(R) Pentium(R) CPU G3250 @ 3.20GHz
BenchmarkGenericReduce-2                49898270                23.70 ns/op
BenchmarkNonGenericReduce-2             154821860                7.527 ns/op
BenchmarkGenericReduceF1-2              50696623                23.51 ns/op
BenchmarkNonGenericReduceF1-2           153482284                7.588 ns/op
BenchmarkGenericReduceF2-2              50168204                23.53 ns/op
BenchmarkNonGenericReduceF2-2           47963671                23.78 ns/op
BenchmarkGenericReduceF3-2              50880426                23.79 ns/op
BenchmarkNonGenericReduceF3-2           50745277                23.51 ns/op
PASS
ok      genericbenchmark        15.145s
```
