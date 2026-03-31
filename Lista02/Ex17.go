package main

import (
	"fmt"
	"strings"
)

func main() {
	var tipoConsumidor string
	var consumo float64
	var total float64

	fmt.Println("--- Sistema de Faturamento SANEAGO ---")
	
	fmt.Print("Digite o tipo do consumidor (Residencial, Comercial, Industrial): ")
	fmt.Scan(&tipoConsumidor)
	tipoConsumidor = strings.ToLower(tipoConsumidor)

	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	switch tipoConsumidor {
	case "residencial":
		total = 5.00 + (consumo * 0.05)

	case "comercial":
		total = 500.00
		if consumo > 80 {
			total += (consumo - 80) * 0.25
		}

	case "industrial":
		total = 800.00
		if consumo > 100 {
			total += (consumo - 100) * 0.04
		}

	default:
		fmt.Println("Erro: Tipo de consumidor inválido.")
		return
	}

	fmt.Printf("\n--- Resumo da Conta ---")
	fmt.Printf("\nTipo: %s", strings.Title(tipoConsumidor))
	fmt.Printf("\nConsumo: %.2f m³", consumo)
	fmt.Printf("\nValor Total: R$ %.2f\n", total)
}