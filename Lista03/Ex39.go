package main

import "fmt"

func main() {
	var peso, maiorPeso, menorPeso float64
	var numero, idMaiorPeso, idMenorPeso, qtd int

	fmt.Print("Quantidade de Bois: ")
	fmt.Scan(&qtd)

	for i := 1; i <= qtd; i++ {

		fmt.Printf("Digite o Indentificador do %dº Boi: ", i)
		fmt.Scan(&numero)

		fmt.Printf("Digite o Peso do %dº Boi: ", i)
		fmt.Scan(&peso)
		if peso > maiorPeso {
			maiorPeso = peso
			idMaiorPeso = numero
		}
		if i == 1 {
			menorPeso = peso
		}
		if peso < menorPeso {
			menorPeso = peso
			idMenorPeso = numero
		}

		fmt.Println("============================================")
	}

	fmt.Println("Identificador: ", idMaiorPeso)
	fmt.Println("Maior peso: ", maiorPeso)
	fmt.Println("============================================")
	fmt.Println("Identificador: ", idMenorPeso)
	fmt.Println("Menor peso: ", menorPeso)
}
