package main

import "fmt"

func PrintSlice(s string, x []int) {
	fmt.Printf("%s = length = %d cap = %d  x = %d\n", s, len(x), cap(x), x)
}
func main() {
	a := make([]int, 5)
	PrintSlice("a", a)
	b := make([]int, 4)
	PrintSlice("b", b)
}
