package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) {
	// two ways to dereference a struct 
	// (*p).first = "Alice"
	p.first = "Alice"
}

func main() {
	p := person{
		first: "Bob",
	}
	fmt.Println(p.first)
	changeMe(&p)
	fmt.Println(p.first)
}
