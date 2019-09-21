package main

import "fmt"

func main() {
	a := increment()
	fmt.Println("A", a())
	fmt.Println("A", a())
	fmt.Println("A", a())
	fmt.Println("A", a())
	a = increment()
	fmt.Println("A2", a())
	fmt.Println("A2", a())
	fmt.Println("A2", a())
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
