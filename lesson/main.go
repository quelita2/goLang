/*   Exercício - GoLang
    Será que passei?

Fazer a média de um aluno e dizer se o mesmo foi aprovado ou vai precisar de recuperação.
*/

package main

import "fmt"

func main() {
	n := 3
	x, err := intScanln(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := somaTudo(n, x...)

	fmt.Println("A média é:", result, ". Por isso...")
	if result < 5 {
		fmt.Print("\nOi recuperação!")
		//rec
	} else if result == 5 {
		fmt.Print("\nPassou raspando, guri")
	} else {
		fmt.Print("\nParabés!")
	}
}

func intScanln(n int) ([]int, error) {
	/*A função make:
	  "make" aloca e inicializa um objeto do tipo slice, map ou chan. Assim como new, o primeiro argumento é um tipo. Mas também pode levar um segundo argumento, o tamanho. Ao contrário de new, o tipo de retorno de make é o mesmo que o tipo de seu argumento, não um ponteiro para ele. E o valor alocado é inicializado (não definido como valor zero como em novo). A razão é que slice, map e chan são estruturas de dados. Eles precisam ser inicializados, caso contrário, não poderão ser usados.
	*/
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}

// Soma as notas e faz a media
func somaTudo(tam int, x ...int) int {
	soma := 0
	for _, i := range x {
		soma += i
	}
	result := soma / tam
	return result
}
