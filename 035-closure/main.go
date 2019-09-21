package main

import "fmt"

func main() {
	a := increment()
	b := increment()
	fmt.Println("A", a())
	fmt.Println("A", a())
	fmt.Println("A", a())
	fmt.Println("A", a())
	fmt.Println("B", b())
	fmt.Println("B", b())
	fmt.Println("B", b())
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
