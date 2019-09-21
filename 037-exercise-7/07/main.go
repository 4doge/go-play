package main

import "fmt"

func main() {
	f := sayHello
	f()

}

func sayHello() {
	fmt.Println("Hello")
}
