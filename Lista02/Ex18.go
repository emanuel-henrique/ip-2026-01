package main

import (
	"fmt"
)

func main() {
	var precoNormal float64
	var diaSemana int
	var categoria int

	fmt.Println("=== SISTEMA DE LOCAÇÃO DE DVDS ===")

	fmt.Print("Digite o preço normal do DVD (R$): ")
	fmt.Scan(&precoNormal)

	fmt.Println("\nDias da semana:")
	fmt.Println("1-Dom, 2-Seg, 3-Ter, 4-Qua, 5-Qui, 6-Sex, 7-Sab")
	fmt.Print("Informe o dia da semana: ")
	fmt.Scan(&diaSemana)

	fmt.Println("\nCategorias:")
	fmt.Println("1 - Comum")
	fmt.Println("2 - Lançamento")
	fmt.Print("Informe a categoria: ")
	fmt.Scan(&categoria)

	fatorDia := 1.0
	switch diaSemana {
	case 2, 3, 5:
		fatorDia = 0.60
		fmt.Println("\n[!] Desconto de 40% aplicado (Seg, Ter ou Qui).")
	case 1, 4, 6, 7:
		fatorDia = 1.0
	default:
		fmt.Println("\n[Erro] Dia da semana inválido.")
		return
	}

	fatorCategoria := 1.0
	if categoria == 2 {
		fatorCategoria = 1.15
		fmt.Println("[!] Acréscimo de 15% aplicado (Lançamento).")
	} else if categoria != 1 {
		fmt.Println("\n[Erro] Categoria inválida.")
		return
	}

	precoFinal := precoNormal * fatorDia * fatorCategoria

	fmt.Println("----------------------------------")
	fmt.Printf("Preço Normal: R$ %.2f\n", precoNormal)
	fmt.Printf("PREÇO FINAL:  R$ %.2f\n", precoFinal)
	fmt.Println("----------------------------------")
}