# Variáveis

## Declaração de variáveis

No Go, há dois modos de declarar uma variável.
  - Primeiro modo onde declaramos com a palavra reservada *var* e passamos o tipo da variável.
  ``` Go
  func main() {
	  var test string = "Hello World!"

    fmt.Println(test) //Retornará o valor: Hello World!
  }
  ```
  - Segundo modo e mais utilizado entre os desenvolvedores, é a declaração através do *:=*, onde o Go interpreta automaticamente o tipo da variável, pelo valor que ela está sendo instanciada.

  ``` Go
  func main() {
	  test := "Hello World!"

	  fmt.Println(test) //Retornará o valor: Hello World!
  }
  ```

## Funções/Variáveis Públicas x Privadas

- Para uma função/variável ser publica a primeira letra dela deve ser maiúscula. Exemplo:

``` Go
  ResultadoSoma := soma(1, 2)

  func Soma(a int, b int) int {
	  return a + b
  }
```

- Para uma função/variável ser privada a primeira letra dela deve ser minúscula. Exemplo:

``` Go
  resultadoSoma := soma(1, 2)

  func soma(a int, b int) int {
	  return a + b
  }
```

## Tipo any no GO

Em outros tipos de linguagens como [**Typescript** e **Kotlin**] quando não sabemos o tipo de retorno, usamos o tipo any. No GO declaramos que uma variável é do tipo any com *interface{}*.

```GO
func main() {
  var test interface{}

  fmt.Printf("%T", test)
}
```

Um outro exemplo seria declarar um JSON que pode receber valores desconhecidos, ficando da seguinte forma:

```GO
func main() {
  testJson := map[string]interface{}{
    "nome": "Aspas",
    "idaide": 19,
  }
}
```