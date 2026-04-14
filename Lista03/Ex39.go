package main

import "fmt"

func main() {
	var peso, maiorPeso float64
	var numero, idMaiorPeso, qtd int

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
		fmt.Println("============================================")
	}

	fmt.Println("Identificador: ", idMaiorPeso)
	fmt.Print("Maior peso: ", maiorPeso)
}
