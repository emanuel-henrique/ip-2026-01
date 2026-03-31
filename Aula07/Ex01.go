package main

import f "fmt"

func main () {

	var notas [3]float64
	var i int
	var soma float64

	f.Println("Insira a primeira nota:")
	f.Scan(&notas[0])
	f.Println("Insira a segunda nota:")
	f.Scan(&notas[1])
	f.Println("Insira a terceira nota:")
	f.Scan(&notas[2])

	for i = 0; i < 3; i++ {
       soma = soma + notas[i]
	}
	media := soma / 3

	f.Printf("Média do aluno = %.2f",media)
	f.Println("")

	for i = 0; i < 3; i++ {
		if notas[i] > media {
			f.Println("A", i + 1,"º nota =", notas[i], ", está acima da média")
		}
	}
}
