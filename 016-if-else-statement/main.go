package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true")
	}

	if false {
		fmt.Println("This is false")
	}

	if !true {
		fmt.Println("This is also false")
	}

	if !false {
		fmt.Println("This is also true")
	}

	if 2 == 2 {
		fmt.Println("This is also also true")
	}

	a := 13
	if a == 10 {
		fmt.Println("a is equal 10")
	} else if a > 10 {
		fmt.Println("a is greater than 10")
	} else {
		fmt.Println("a is less than 10")
	}
}
