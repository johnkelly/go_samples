package main

import "fmt"

func main() {
	num := IterativeFib(10)
	fmt.Println(num)
}

//0,1,1,2,3,5,8,13,21,34
func IterativeFib(n int) int {
	x := 0
	y := 1

	for i := 0; i < n; i++ {
		x, y = y, (x + y)
	}

	return x
}

func RecursiveFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return RecursiveFib(n-1) + RecursiveFib(n-2)
	}
}

//Results to calculate first 10
// BenchmarkIterativeFib-8	200000000	         6.76 ns/op
// BenchmarkRecursiveFib-8	 5000000	       333 ns/op

//Results to calculate first 30
// BenchmarkIterativeFib-8	100000000	        15.5 ns/op
// BenchmarkRecursiveFib-8	     300	   4940570 ns/op

//Results to calculate first 70
// BenchmarkIterativeFib-8	30000000	        41.2 ns/op
// BenchmarkRecursiveFib-8	SIGQUIT: quit
