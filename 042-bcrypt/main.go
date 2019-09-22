package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "testpass"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(bs, []byte("anotherpassword123"))
	if err != nil {
		fmt.Println(err)
	}
	
}
