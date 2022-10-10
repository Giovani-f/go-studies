package main

import "fmt"

func main() {
	var user User = User{
		Name:     "Neymar",
		Age:      30,
		password: "teste123",
	}

	fmt.Println(user)
}

type User struct {
	Name     string // campo público
	Age      int    // campo público
	password string // campo privado
}
