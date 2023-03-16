package main // O Go está relaciado a pacotes, e para iniciar o código, vamos usar o pacote 'main' de menu principal (função pincipal do programa, que é o ponto de entrada)

import ( // importa pacotes ao invés de importar bibliotecas externas
	"fmt" // pacote padrão: 'format'
  "learningGo/math" // importando um pacote criado
)

var a string // declaração de variável

func main() {
	fmt.Println("\nHello, World!")
  a = "Quelita" //atribuição da variável
  // Para declarar (:) e atribuir (=) valor a variável, basta aplicar o ':=' que a linguagem adivinha o tipo da variável.
  b := 20 // o Go n compila uma variável que não está                           sendo utilizada!!!
  c := 2.34
  d := false

  // Mostrar os valores atribuídos das variáveis
  fmt.Printf("\n%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n", d)

  // Mostrar o type das variáveis
  fmt.Printf("\n%T \n", a)
  fmt.Printf("%T \n", b)
  fmt.Printf("%T \n", c)
  fmt.Printf("%T \n", d)

  // resultado := math.Soma(1,1)
  // fmt.Printf("A soma é: %v", resultado)
  fmt.Print("A soma é: ",math.Soma(1,1))
}