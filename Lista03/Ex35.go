package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo na base 10: ")
	fmt.Scan(&n)

	fmt.Printf("O número %d em binário (base 2) é: %b\n", n, n)
}
