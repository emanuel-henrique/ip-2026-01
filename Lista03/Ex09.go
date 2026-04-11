package main

import "fmt"

func main() {
	var n, aprovados, exame, reprovados int
	var n1, n2, somaClasse float64

	fmt.Print("Quantidade de alunos (N): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Notas do aluno %d: ", i)
		fmt.Scan(&n1, &n2)

		media := (n1 + n2) / 2
		somaClasse += media

		fmt.Printf("Média: %.2f - ", media)
		if media >= 7 {
			fmt.Println("Aprovado")
			aprovados++
		} else if media >= 3 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Reprovado")
			reprovados++
		}
	}

	fmt.Printf("\nTotal Aprovados: %d", aprovados)
	fmt.Printf("\nTotal Exame: %d", exame)
	fmt.Printf("\nTotal Reprovados: %d", reprovados)
	fmt.Printf("\nMédia da Classe: %.2f\n", somaClasse/float64(n))
}