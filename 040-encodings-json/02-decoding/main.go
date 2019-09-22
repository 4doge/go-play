package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string `json:"Name"`
	Occupation string `json:"Occupation"`
	Age        int    `json:"Age"`
}

func main() {
	s := `[{"Name":"Bob","Occupation":"IT guy","Age":27},{"Name":"Alice","Occupation":"Instagirl","Age":24}]`
	bs := []byte(s)
	people := []person{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v.Name)
	}
}
