package main

import "fmt"

var a int
type mycustomtype int
var b mycustomtype

func main() {
	a = 12
	b = 13
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}