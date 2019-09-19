package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("not printing")
	case (2+2 == 4):
		fmt.Println("print")
		fallthrough
	case (3 == 4):
		fmt.Println("not printing 2")
		fallthrough
	case (3 == 3):
		fmt.Println("print 2")
		fallthrough
	default:
		fmt.Println("default case")
	}

	switch "example" {
	case "exa":
		fmt.Println("exa")
	case "e", "x", "a":
		fmt.Println("e or x or a")
	case "ex", "example":
		fmt.Println("Example!")
	default:
		fmt.Println("Default")
	}
}
