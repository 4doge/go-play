package main

import "fmt"

type person struct {
	name       string
	occupation string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hey. I'm", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	b := person{
		name: "Bob",
		occupation: "IT guy",
	}
	saySomething(&b)
	// Below is not working code because of the passing the value instead of a pointer
	// saySomething(b)
}
