package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	b := person{
		firstName: "Bob",
		lastName:  "Smith",
		favoriteIceCream: []string{
			"Chocolate",
			"Hazelnut",
		},
	}

	a := person{
		firstName: "Alice",
		lastName:  "Smith",
		favoriteIceCream: []string{
			"Strawberry",
			"Vanilla",
		},
	}

	fmt.Println(a, b)

	for i, v := range a.favoriteIceCream {
		fmt.Println(i, v)
	}

}
