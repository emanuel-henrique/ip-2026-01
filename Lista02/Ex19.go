package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var r, h float64

	fmt.Println("=== CÁLCULO DE SÓLIDOS GEOMÉTRICOS ===")
	fmt.Println("1 - Cone Reto")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Print("Escolha uma opção: ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Print("Digite o raio (r): ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura (h): ")
		fmt.Scan(&h)

		g := math.Sqrt(math.Pow(r, 2) + math.Pow(h, 2))
		volume := (math.Pi * math.Pow(r, 2) * h) / 3
		area := math.Pi * r * (r + g)

		fmt.Printf("\n--- Cone ---\nVolume: %.2f\nÁrea da Superfície: %.2f\n", volume, area)

	case 2:
		fmt.Print("Digite o raio (r): ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura (h): ")
		fmt.Scan(&h)

		volume := math.Pi * math.Pow(r, 2) * h
		area := 2 * math.Pi * r * (r + h)

		fmt.Printf("\n--- Cilindro ---\nVolume: %.2f\nÁrea da Superfície: %.2f\n", volume, area)

	case 3:
		fmt.Print("Digite o raio (r): ")
		fmt.Scan(&r)

		volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		area := 4 * math.Pi * math.Pow(r, 2)

		fmt.Printf("\n--- Esfera ---\nVolume: %.2f\nÁrea da Superfície: %.2f\n", volume, area)

	default:
		fmt.Println("Opção inválida!")
	}
}