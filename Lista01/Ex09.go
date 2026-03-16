package main

import "fmt"

// Cálculo do Determinante de uma Matriz Quadrada de Duas Dimensões
func main() {
	var a, b, c, d float64

	fmt.Print("Digite o valor do elemento 'A': ")
	fmt.Scan(&a)

	fmt.Print("Digite o valor do elemento 'B': ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor do elemento 'C': ")
	fmt.Scan(&c)

	fmt.Print("Digite o valor do elemento 'D': ")
	fmt.Scan(&d)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", calculateDeterminant(a,b,c,d))
}

func calculateDeterminant(a, b, c, d float64) float64 {
	determinant := a*d - b*c

	return determinant
}