package main

import "fmt"

func main() {
	f := foo(bar)
	fmt.Println(f)
}

func foo(f func() int) int {
	fmt.Println("Hey from foo")
	return f()
}

func bar() int {
	fmt.Println("Hey from bar")
	return 13
}
