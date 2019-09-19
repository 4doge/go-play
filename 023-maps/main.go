package main

import "fmt"

func main() {
	m := map[string]int{
		"bob": 27,
		"alice": 23,
	}
	fmt.Println(m)
	fmt.Println(m["bob"])
	fmt.Println(m["example"])

	v, ok := m["example"]
	fmt.Println(v, ok)
	if v, ok := m["example"]; ok {
		fmt.Println("Exist", v)
	}
	

	m["eve"] = 25

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "eve")

	fmt.Println(m)
}
