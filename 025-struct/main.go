package main

import "fmt"

type person struct {
	name       string
	occupation string
	age        int
}

type superman struct {
	person
	abilityToFly bool
}

func main() {
	b := person{
		name:       "Bob",
		occupation: "IT guy",
		age:        27,
	}
	a := person{
		name:       "Alice",
		occupation: "Instagirl",
		age:        24,
	}
	fmt.Printf("%T\n", a)
	fmt.Println(a, b)

	c := superman{
		person: person{
			name:       "Clark",
			occupation: "Journalist",
			age:        123,
		},
		abilityToFly: true,
	}
	fmt.Printf("%T\n %v\n", c, c)
	fmt.Println(c.name, c.occupation)

	// anonymous struct
	d := struct {
		name    string
		surname string
	}{
		name:    "John",
		surname: "Doe",
	}
	fmt.Println(d)
}
