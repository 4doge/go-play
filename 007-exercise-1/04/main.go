package main

import "fmt"

type mycustomtype int

var x mycustomtype

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)
}