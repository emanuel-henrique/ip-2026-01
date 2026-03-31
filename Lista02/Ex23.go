package main

import (
	"fmt"
)

func main() {
	var idade int

	fmt.Println("=== CONSULTA DE CLASSE ELEITORAL ===")
	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	switch {
	case idade < 16:
		fmt.Println("Classe: Não-eleitor")

	case (idade >= 16 && idade < 18) || idade > 65:
		fmt.Println("Classe: Eleitor facultativo")

	case idade >= 18 && idade <= 65:
		fmt.Println("Classe: Eleitor obrigatório")

	default:
		fmt.Println("Idade inválida.")
	}
}