package main

import "fmt"

type mycustomtype int
var x mycustomtype
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 13
	fmt.Printf("%v %T\n", x, x)
	y = int(x)
	fmt.Printf("%v %T\n", y, y)
}