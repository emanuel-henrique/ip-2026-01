package main

import (
	"fmt"
)

func mdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2 inteiros positivos: ")
	fmt.Scan(&n1, &n2)

	resultado := mmc(n1, n2)
	fmt.Printf("O MMC de %d e %d é: %d\n", n1, n2, resultado)
}
