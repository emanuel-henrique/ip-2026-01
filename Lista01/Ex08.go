package main

import "fmt"

// Cálculo do Delta na Equação de Báskara
func main() {
	var a,b,c float64

	fmt.Print("Digite o valor de 'A': ")
	fmt.Scan(&a)

	fmt.Print("Digite o valor de 'B': ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor de 'C': ")
	fmt.Scan(&c)

	fmt.Printf("O VALOR DE DELTA E = %.2f", calculateDelta(a,b,c))
}

func calculateDelta(a, b, c float64) float64 {
	var delta = (b * b) - 4*a*c
	return delta
}