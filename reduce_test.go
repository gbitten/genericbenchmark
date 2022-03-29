package main

import (
	"testing"
)

func BenchmarkGenericReduce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenericReduce[int,int](Numbers, func(acc, current int) int { return acc + current }, 0)
	}
}

func BenchmarkNonGenericReduce1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NonGenericReduce(Numbers, func(acc, current int) int { return acc + current }, 0)
	}
}

func BenchmarkGenericReduce2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenericReduce[int,int](Numbers, F, 0)
	}
}

func BenchmarkNonGenericReduce2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NonGenericReduce(Numbers, F, 0)
	}
}