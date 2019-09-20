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

func (s superman) hateBatman() {
	fmt.Println("Hate Batman! Because I'm the", s.name, "- superman")
}

func main() {
	bobTheSuperman := superman{
		person: person{
			name:       "Bob",
			occupation: "IT guy",
			age:        27,
		},
		abilityToFly: true,
	}
	aliceTheSuperman := superman{
		person: person{
			name:       "Alice",
			occupation: "Instagirl",
			age:        24,
		},
		abilityToFly: false,
	}

	fmt.Println(bobTheSuperman)
	bobTheSuperman.hateBatman()
	aliceTheSuperman.hateBatman()
}
