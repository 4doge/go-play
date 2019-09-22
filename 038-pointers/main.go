package main

import "fmt"

func main() {
	a := "Hello world"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a)
	fmt.Printf("%T\n\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
	fmt.Printf("%T\n", *b)
}
