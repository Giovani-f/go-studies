# Estruturas de controle

## IF e ELSE

Diferente de algumas linguagens no Go não usamos parentese envolta da condição do if.
  - Estrutra de controle IF:
```Go
func main() {
  if 2 > 1 {
    fmt.Print("Hello GO!")
  }
}
```
  - Estrutra de controle IF e ELSE:
  - Um detalhe no else é que ele deve sempre seguir após o fechamento do if, caso quebre uma linha o código não irá compilar.
```Go
func main() {
  if 2 > 1 {
    fmt.Print("Hello GO!")
  } else {
    fmt.Print("Hello world!")
  }
}
```

  - Estrutra de controle ELSE IF:
```Go
func main() {
  if 2 > 1 {
    fmt.Print("Hello GO!")
  } else if 2 > 3 {
    fmt.Print("Hello world!")
  }
}
```

## IF com inicialização

No GO podemos criar uma váriavel diretamente na condição do IF para usuarmos em uma validação. Exemplo GO:

```Go
func main() {
  if result := test(); result == "test" {
    fmt.Print(true)
  }
}

func test() string {
  return "test"
}
```

- Exemplo do mesmo código feito em Javascript:

```javascript
const result = test()
if (restul === "test") {
  console.log(true)
}

function test() {
  return "test"
}
```

- Essa váriavel que é inicializada na condição do if so irá existir no escopo do if.

## Switch

Exemplo de código com a estrutura de controle switch:

```Go
func main() {
  conditionSwitch := 1

  switch conditionSwitch {
  case 1:
    fmt.Println("Caiu na primeira codição!")

  case 2:
    fmt.Println("Caiu na segunda codição!")
  }
}
```

- O *case* pode aceitar mais de um valor, como no exemplo a seguir:

```Go
func main() {
  conditionSwitch := 1

  switch conditionSwitch {
  case 1, 3:
    fmt.Println("Caiu na primeira codição!")

  case 2:
    fmt.Println("Caiu na segunda codição!")
  }
}
```

### Fallthrough

Em outras linguagens quando temos um switch, adicionamos a palavra reservada **break** para que a estrutura não continue executando mesmo após passar por sua condição valida.

No Go isso funciona ao contrário, por padrão ao atingir a condição ele não irá executar os próximos *cases*. Caso queira que isso aconteça utilizamos da palavra reservada **Fallthrough**, para que execute o case a seguir, mesmo que a condição não seja válida.

```Go
func main() {
  conditionSwitch := 1

  switch conditionSwitch {
  case 1:
    fmt.Println("Caiu na primeira codição!")
    fallthrough

  case 2: // irá executar, pois na condição acima possuimos o fallthrough
    fmt.Println("Caiu na segunda codição!")

  case 3: // mão irá executar, pois na condição acima não possuimos o fallthrough
    fmt.Println("Caiu na terceira codição!")
  }
}
```

### Default

Exemplo de código com a condição default:

```Go
func main() {
  conditionSwitch := 1

  switch conditionSwitch {
  case 1:
    fmt.Println("Caiu na primeira codição!")
    fallthrough

  case 2:
    fmt.Println("Caiu na segunda codição!")

  default:
    fmt.Println("Caiu no default!")
  }
}
```