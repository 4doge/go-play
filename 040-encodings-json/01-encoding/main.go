package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string
	Occupation string
	Age        int
}

func main() {
	bob := person{
		Occupation: "IT guy",
		Age:        27,
	}
	alice := person{
		Name:       "Alice",
		Occupation: "Instagirl",
		Age:        24,
	}

	people := []person{bob, alice}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
