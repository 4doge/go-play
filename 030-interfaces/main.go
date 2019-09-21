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

func (p person) hateBatman() {
	fmt.Println("Hate Batman! Because I'm the", p.name, "- regular guy")
}

type hero interface {
	hateBatman()
}

func heroScream(h hero) {
	switch h.(type) {
	case person:
		fmt.Println("Yaaa yeee", h.(person).name)
	case superman:
		fmt.Println("Raaaaaaagr", h.(superman).name)
	}
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

	regularGuy := person{
		name:       "John",
		occupation: "do nothing",
		age:        44,
	}

	fmt.Println(regularGuy)
	regularGuy.hateBatman()

	heroScream(bobTheSuperman)
	heroScream(aliceTheSuperman)
	heroScream(regularGuy)

}
