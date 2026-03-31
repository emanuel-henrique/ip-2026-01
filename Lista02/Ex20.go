package main

import (
	"fmt"
)

func main() {
	var precoEtiqueta float64
	var codigo int
	var valorFinal float64

	fmt.Println("=== SISTEMA DE PAGAMENTO ===")
	
	fmt.Print("Digite o preço de etiqueta do produto (R$): ")
	fmt.Scan(&precoEtiqueta)

	fmt.Println("\nCondições de Pagamento:")
	fmt.Println("1 - À vista (Dinheiro/Cheque): 10% de desconto")
	fmt.Println("2 - À vista (Cartão de Crédito): 5% de desconto")
	fmt.Println("3 - Em 2x: Preço normal (sem juros)")
	fmt.Println("4 - Em 3x: Preço normal + 10% de juros")
	fmt.Print("\nEscolha o código da condição: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 1:
		valorFinal = precoEtiqueta * 0.90
		fmt.Printf("\nDesconto aplicado: 10%%")

	case 2:
		valorFinal = precoEtiqueta * 0.95
		fmt.Printf("\nDesconto aplicado: 5%%")

	case 3:
		valorFinal = precoEtiqueta
		fmt.Printf("\nPagamento em 2 parcelas de: R$ %.2f", valorFinal/2)

	case 4:
		valorFinal = precoEtiqueta * 1.10
		fmt.Printf("\nJuros aplicados: 10%%")
		fmt.Printf("\nPagamento em 3 parcelas de: R$ %.2f", valorFinal/3)

	default:
		fmt.Println("\n[Erro] Código de pagamento inválido!")
		return
	}

	fmt.Println("\n--------------------------------")
	fmt.Printf("Valor Total a Pagar: R$ %.2f\n", valorFinal)
	fmt.Println("--------------------------------")
}