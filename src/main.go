package main

import "fmt"

// a: the value of the n-2 ths fib number a >= 0
// b: the value of the n-1 th fib number. b >= 1
func fib(n, a, b int) int {
	if n == 0 {
		return a
	}

	if n == 1 {
		return b
	}

	return fib(n-1, b, a+b)
}

func main() {
	result := fib(10, 0, 1)
	fmt.Println(result) // should be 55
}
