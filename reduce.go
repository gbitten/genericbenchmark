package main

import (
    "fmt"
)

var Numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func F1 (acc, current int) int { return acc + current }
var F2 = F1
var F3 = func(acc, current int) int { return acc + current }


func GenericReduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func NonGenericReduce ( s[] int, f func(int, int) int, initValue int) int {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func main() {
    sum := GenericReduce(Numbers, F1, 0)
    fmt.Println(sum)
    sum = NonGenericReduce(Numbers, F1, 0)
    fmt.Println(sum)
    sum = GenericReduce(Numbers, F2, 0)
    fmt.Println(sum)
    sum = NonGenericReduce(Numbers, F2, 0)
    fmt.Println(sum)
    sum = GenericReduce(Numbers, F3, 0)
    fmt.Println(sum)
    sum = NonGenericReduce(Numbers, F3, 0)
    fmt.Println(sum)
}
