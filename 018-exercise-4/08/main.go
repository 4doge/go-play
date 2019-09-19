package main

import "fmt"

func main() {
	switch {
	case 13 == 14:
		fmt.Println("not equal")
	case 2*2 == 4:
		fmt.Println("this is true")
	}
}
