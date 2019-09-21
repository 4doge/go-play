package main

import "fmt"

func main() {
	helloAlice := sayHello("Alice")
	helloBob := sayHello("Bob")
	fmt.Println(helloAlice)
	fmt.Println(helloBob)
	f := generateNumber()
	fmt.Printf("%T\n", f)
	num := f()
	fmt.Printf("%T %v\n", num, num)
	fmt.Println(generateNumber()())
}

func sayHello(name string) string {
	return "Hello, " + name
}

func generateNumber() func() int {
	return func() int {
		return 13
	}
}
