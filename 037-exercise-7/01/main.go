package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a)
	fmt.Println(b, c)
}

func foo() int {
	return 13
}

func bar() (int, string) {
	return 404, "Hey"
}
