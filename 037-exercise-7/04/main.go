package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hey, My name is", p.first, p.last, "and my age is", p.age)
}

func main() {
	p := person{
		first: "Bob",
		last:  "Smith",
		age:   27,
	}
	p.speak()
}
