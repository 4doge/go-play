package main

import "fmt"

func main() {
	func(who string) {
		fmt.Println("Just to say hello from ", who)
	}("anonymous function")
}
