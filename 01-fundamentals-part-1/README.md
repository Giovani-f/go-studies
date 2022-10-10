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

## Arrays/Slices

A diferença entre array e slice é que um Array possui um tamanho definido, já um slice possui um tamanho indefinido.

### Definindo um Array

```GO
func main() {
  var testArray [4]string = [4]string{"Carro", "Moto", "Bicicleta", "Patins"}
  fmt.Println(len(testArray))
  fmt.Println(cap(testArray)) // mostra a capacidade de um slice/array
}
```

### Definindo um Slice

O slice irá crescendo conforme a quantidade de valores que forem colocados nele.

```GO
func main() {
  var testArray []string = []string{"Carro", "Moto", "Bicicleta", "Patins", "Skate"}
  fmt.Println(len(testArray))
  fmt.Println(cap(testArray)) // mostra a capacidade de um slice/array
}
```

## Structs

Structs em Go são como *Objetos* em outras liguagens como no *Typescript* e são definidos da seguinte forma:

### Definindo uma struct

```GO
type User struct {
  name string
  age int
}
```

### Utilizando uma struct

```GO
func main() {
  var user User = User{
    name: "Neymar",
    age: 30,
  }

  fmt.Print(user)
}

type User struct {
  name string
  age int
}
```