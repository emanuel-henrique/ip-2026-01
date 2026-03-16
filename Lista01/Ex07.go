package main

import (
	"fmt"
	"math"
)

// Custo da Lata de Cerveja

var pi float64 = math.Pi 

func main() {
	var r, h float64

	fmt.Print("Digite o valor do raio da lata: ")
	fmt.Scan(&r)
	fmt.Print("Digite o valor da altura da lata: ")
	fmt.Scan(&h)

	fmt.Printf("O VALOR DO CUSTO E = %.2f", calcularCustoLata(r,h))
}

func calcularCustoLata(r float64, h float64) float64 {
	areaBase := 2 * (pi * (r * r)) + (2 * pi * r * h)
	return areaBase * 100
}

