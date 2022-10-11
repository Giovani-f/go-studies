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

	//switch
	conditionSwitch := 1

	switch conditionSwitch {
	case 1, 3:
		fmt.Println("Caiu na primeira codição!")

	case 2:
		fmt.Println("Caiu na segunda codição!")
	}

	// switch como fallthrough
	switch conditionSwitch {
	case 1:
		fmt.Println("Caiu na primeira codição fallthrough!")
		fallthrough

	case 2:
		fmt.Println("Caiu na segunda codição fallthrough!")

	case 3:
		fmt.Println("Caiu na terceira codição fallthrough!")
	}

	// switch com default
	conditionSwitchDefault := 3
	switch conditionSwitchDefault {
	case 1:
		fmt.Println("Caiu na primeira codição!")

	case 2:
		fmt.Println("Caiu na segunda codição!")

	default:
		fmt.Println("Caiu no default!")
	}
}

func ifWithInit() string {
	return "test"
}
