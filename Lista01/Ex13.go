package main

import (
	"fmt"
	"math"
)

// Volume da Pirâmide de Base Hexagonal
func main() {
	var a, l float64

	fmt.Print("Digite a altura da pirâmide: ")
	fmt.Scan(&a)

	fmt.Print("Digite o valor da aresta da base da pirâmide: ")
	fmt.Scan(&l)

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", calculateVolume(l, a))
}

func calculateHexagonArea(a float64) float64{
	area := (3 * (a*a) * math.Sqrt(3)) / 2
	return area
}

func calculateVolume(l, h float64) float64{
	aBase := calculateHexagonArea(l)
	volume := (aBase * h) / 3
	return volume
}
