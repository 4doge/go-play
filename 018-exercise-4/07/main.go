package main

import "fmt"

func main() {
	x := 15
	if x > 13 {
		fmt.Println("x is greater than 13")
	} else if x < 13 {
		fmt.Println("x is less than 13")
	} else {
		fmt.Println("x is 13")
	}
}
