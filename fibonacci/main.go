package main

import "fmt"

func fibonacci(n int) {
	fib1 := 0
	fib2 := 1
	for i := 0; i < n; i++ {
		fib3 := fib1 + fib2
		fib1 = fib2
		fib2 = fib3
		fmt.Printf("%d\t", fib3)
	}
}

func main() {
	fibonacci(10)
}
