package main

import "fmt"

func isPalindome(s string) bool {
	b:=len(s)
	for	i:=0;i<=len(s)/2;i++ {
		if(s[i]!=s[b-i-1]) {
		    return false
		}
	}
	return true
}

func main() {
	a:=isPalindome("lool")
	if(a) {
		fmt.Println("It is palindrome")
	} else {
		fmt.Println("It is not palindrome")
	}
}

