package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("FOR", i)
	}

	// while
	test := 0

	for test <= 3 {
		fmt.Println("WHILE", test)
		test++
	}

	// do while
	expression := false
	for ok := true; ok; ok = expression {
		fmt.Println("DO-WHILE")
	}

	//for range
	testRange := []string{"range1", "range2", "range3"}

	for i, value := range testRange {
		fmt.Println(value, i)
	}
}
