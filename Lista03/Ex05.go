package main

import "fmt"

func main() {
	var idade, continuar int
	var altura, peso, somaAltura1020, qtdPessoas, qtdAcima50, qtdAltura1020, qtdPesoAbaixo40 float64

	for {
		fmt.Print("Idade: ")
		fmt.Scan(&idade)
		fmt.Print("Altura: ")
		fmt.Scan(&altura)
		fmt.Print("Peso: ")
		fmt.Scan(&peso)

		qtdPessoas++

		if idade > 50 {
			qtdAcima50++
		}
		if idade >= 10 && idade <= 20 {
			somaAltura1020 += altura
			qtdAltura1020++
		}
		if peso < 40 {
			qtdPesoAbaixo40++
		}

		fmt.Print("Deseja continuar? (1-Sim, Outro-Não): ")
		fmt.Scan(&continuar)
		if continuar != 1 {
			break
		}
	}

	fmt.Printf("Pessoas com mais de 50 anos: %.0f\n", qtdAcima50)
	if qtdAltura1020 > 0 {
		fmt.Printf("Média das alturas (10-20 anos): %.2f\n", somaAltura1020/qtdAltura1020)
	}
	fmt.Printf("Porcentagem com peso < 40kg: %.2f%%\n", (qtdPesoAbaixo40/qtdPessoas)*100)
}