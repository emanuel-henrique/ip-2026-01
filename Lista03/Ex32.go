package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2: ")
	fmt.Scan(&n1, &n2)

	negativo := false
	if n1 < 0 && n2 > 0 || n1 > 0 && n2 < 0 {
		negativo = true
	}

	if n1 < 0 { n1 = -n1 }
	if n2 < 0 { n2 = -n2 }

	resultado := 0
	for i := 0; i < n2; i++ {
		resultado += n1
	}

	if negativo {
		resultado = -resultado
	}

	fmt.Printf("A multiplicação é: %d\n", resultado)
}
