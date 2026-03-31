package main

import (
	"fmt"
)

func main() {
	var id int
	var n1, n2, n3, mediaExercicios float64
	var conceito string
	var status string

	fmt.Println("=== SISTEMA DE AVALIAÇÃO ACADÊMICA ===")

	fmt.Print("Número de Identificação do Aluno: ")
	fmt.Scan(&id)
	fmt.Print("Nota 1: ")
	fmt.Scan(&n1)
	fmt.Print("Nota 2: ")
	fmt.Scan(&n2)
	fmt.Print("Nota 3: ")
	fmt.Scan(&n3)
	fmt.Print("Média dos Exercícios: ")
	fmt.Scan(&mediaExercicios)

	mediaFinal := (n1 + (n2 * 2) + (n3 * 3) + mediaExercicios) / 7

	switch {
	case mediaFinal >= 9.0:
		conceito = "A"
		status = "APROVADO"
	case mediaFinal >= 7.5:
		conceito = "B"
		status = "APROVADO"
	case mediaFinal >= 6.0:
		conceito = "C"
		status = "APROVADO"
	case mediaFinal >= 4.0:
		conceito = "D"
		status = "REPROVADO"
	default:
		conceito = "E"
		status = "REPROVADO"
	}

	fmt.Println("\n--------------------------------")
	fmt.Printf("Aluno ID: %d\n", id)
	fmt.Printf("Notas: %.1f, %.1f, %.1f\n", n1, n2, n3)
	fmt.Printf("Média Exercícios: %.1f\n", mediaExercicios)
	fmt.Printf("Média de Aproveitamento: %.2f\n", mediaFinal)
	fmt.Printf("CONCEITO: %s\n", conceito)
	fmt.Printf("STATUS FINAL: %s\n", status)
	fmt.Println("--------------------------------")
}