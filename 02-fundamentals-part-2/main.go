package main

import "fmt"

func main() {
	if 2 > 1 {
		fmt.Println("Hello GO!")
	} else if 2 > 3 {
		fmt.Println("Hello world!")
	} else {
		fmt.Println("Dale")
	}

	// if como inicialzação
	if result := ifWithInit(); result == "test" {
		fmt.Println("IF com init")
	}
}

func ifWithInit() string {
	return "test"
}
