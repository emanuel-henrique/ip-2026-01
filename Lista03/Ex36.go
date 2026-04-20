package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo na base 10: ")
	fmt.Scan(&n)

	fmt.Printf("O número %d em hexadecimal (base 16) é: %X\n", n, n)
}
