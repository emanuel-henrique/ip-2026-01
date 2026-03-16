package main

import (
	"fmt"
)

// Consumo de energia
func main() {
	var salarioMinimo float64
	var KwGastos float64

	fmt.Print("Digite o valor do salário minimo: ")
	fmt.Scan(&salarioMinimo)

	fmt.Print("Digite a quantidade de Kw gastos: ")
	fmt.Scan(&KwGastos)

	custoPorKw := (salarioMinimo * 0.7) / 100
	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKw)
	custoConsumo := custoPorKw * KwGastos
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoConsumo * 0.9)
}