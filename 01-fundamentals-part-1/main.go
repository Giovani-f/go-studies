package main

import "fmt"

func main() {
	var test string = "Hello World!"
	fmt.Println(test)

	test2 := "Hello World!"
	fmt.Println(test2)

	// função/variável pública
	ResultadoSoma := Soma(1, 2)
	fmt.Println(ResultadoSoma)

	// função/variável privada
	resultadoSoma := soma(2, 2)
	fmt.Println(resultadoSoma)

	//tipo any
	var testAny interface{}
	fmt.Printf("%T", testAny)

	//JSON com valores any
	testJson := map[string]interface{}{
		"nome":   "Aspas",
		"idaide": 19,
	}
	fmt.Printf("%T", testJson)

	// declaração de array
	var testArray [4]string = [4]string{"Carro", "Moto", "Bicicleta", "Patins"}
	fmt.Println(len(testArray))
	fmt.Println(cap(testArray))

	// declaração de slice
	var testSlice [4]string = [4]string{"Carro", "Moto", "Bicicleta", "Patins"}
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))
}

func Soma(a int, b int) int {
	return a + b
}

func soma(a int, b int) int {
	return a + b
}
