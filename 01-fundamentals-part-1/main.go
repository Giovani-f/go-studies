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
	fmt.Println(testAny)

	//JSON com valores any
	testJson := map[string]interface{}{
		"nome":   "Aspas",
		"idaide": 19,
	}
	fmt.Println(testJson)

	// declaração de array
	var testArray [4]string = [4]string{"Carro", "Moto", "Bicicleta", "Patins"}
	fmt.Println(len(testArray))
	fmt.Println(cap(testArray))

	// declaração de slice
	var testSlice [4]string = [4]string{"Carro", "Moto", "Bicicleta", "Patins"}
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))
}

// Função publica
func Soma(a int, b int) int {
	return a + b
}

// Função privada
func soma(a int, b int) int {
	return a + b
}
