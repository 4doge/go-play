package main

import "fmt"

func main() {
	s := struct {
		m map[string][]string
		s []string
	}{
		m: map[string][]string{
			"names": []string{
				"Bob",
				"Alice",
			},
		},
		s: []string{
			"whatever",
			"some",
			"text",
		},
	}

	fmt.Println(s)
}
