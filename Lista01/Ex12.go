package main

import "fmt"

// Conversão de Nota em Conceito
func main() {
	var grade float64

	fmt.Print("Digite a nota do aluno: ")
	fmt.Scan(&grade)

	fmt.Printf("NOTA = %.1f CONCEITO = %s", grade, getStudentGrade(grade))
}

func getStudentGrade(grade float64) string {
	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade >= 7.5 && grade < 9:
		return "B"
	case grade >= 6 && grade < 7.5:
		return "C"
	default:
		return "D"
	}
}