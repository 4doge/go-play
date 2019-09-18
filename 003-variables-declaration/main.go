package main

import "fmt"

var x = 10

// define zero based variable
// 0 for int
// false for bool
// 0.0 for float
// "" for string
// nil for pointers / func tions / interfaces / slices / channels / maps
var y int

func main() {
	z := 15
	fmt.Println("x - ", x)
	fmt.Println("y - ", y)
	fmt.Println("z - ", z)
}