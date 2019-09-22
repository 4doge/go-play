package main

import "fmt"

func main() {
	//  & - return the memory cell address
	//  * - return the memory cell value
	a := "Hello world"
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a)
	fmt.Println(*&a)
	fmt.Printf("%T\n\n", &a)


	b := &afmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)
	fmt.Printf("%T\n\n", *b)

	*b = "Changed hello world string"
	fmt.Println(a)

}
