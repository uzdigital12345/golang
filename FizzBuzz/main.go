package main

import "fmt"

func main() {
	for index := 1; index < 100; index++ {
		if index%15 == 0 {
			fmt.Print("Fizz Buzz, ")
		} else if index%3 == 0 {
			fmt.Print("Fizz, ")
		} else if index%5 == 0 {
			fmt.Print("Buzz, ")
		} else {
			fmt.Printf("%d, ", index)
		}
	}
}
