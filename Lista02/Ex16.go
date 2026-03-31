package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Calculadora de Equação do 2º Grau (Ax² + Bx + C = 0)")
	
	fmt.Print("Digite o coeficiente A: ")
	fmt.Scan(&a)

	if a == 0 {
		fmt.Println("O coeficiente A não pode ser zero em uma equação do 2º grau.")
		return
	}

	fmt.Print("Digite o coeficiente B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o coeficiente C: ")
	fmt.Scan(&c)

	delta := math.Pow(b, 2) - (4 * a * c)

	fmt.Printf("\nDelta calculado: %.2f\n", delta)

	switch {
	case delta < 0:
		fmt.Println("Classificação: RAÍZES IMAGINÁRIAS")
		
	case delta == 0:
		x := -b / (2 * a)
		fmt.Println("Classificação: RAIZ ÚNICA")
		fmt.Printf("X = %.2f\n", x)
		
	case delta > 0:
		sqrtDelta := math.Sqrt(delta)
		x1 := (-b + sqrtDelta) / (2 * a)
		x2 := (-b - sqrtDelta) / (2 * a)
		
		fmt.Println("Classificação: RAÍZES DISTINTAS")
		fmt.Printf("X1 = %.2f\n", x1)
		fmt.Printf("X2 = %.2f\n", x2)
	}
}