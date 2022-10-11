# Estruturas de repetição

## FOR

A única diferença para outras linguagens no FOR é que não possui parenteses envolta da condição. Exemplo:

```Go
func main() {
  for i := 0; i < 5; i++ {
    fmt.Println("FOR", i)
  }
}
```

- No Go não possuimos a estrutura de repetição while, porém podemos reproduzi-la da seguinte forma:

```Go
func main() {
	test := 0

	for test <= 10 {
	  fmt.Println("WHILE", test)
      test++
	}
}
```

## DO-WHILE

Exemplo de como seria o do while em GO:

```Go
func main() {
	expression := false

	for ok := true; ok; ok = expression {
	  fmt.Println("DO-WHILE", test)
	}
}
```

## FOR com Range

O for com range irá fornecer 2 variáveis, o indice sendo o primeiro valor e o valor em si. Exemplo:

```Go
func main() {
	test := []string{ "test1", "test2", "test3", }

  for i, value := range test {
    fmt.Println(value, i)
  }
}
```

- Caso o indice não seja relevante ao seu escopo de código, basta colocar um **_**, o undelerline irá ignorar o retorno que virá. Exemplo:

```Go
func main() {
	test := []string{ "test1", "test2", "test3", }

  for _, value := range test {
    fmt.Println(value)
  }
}
```
