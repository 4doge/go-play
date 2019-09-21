package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("We need to ge deeper")
	}()
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}
