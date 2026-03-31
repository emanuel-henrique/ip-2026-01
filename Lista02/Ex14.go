package main

import "fmt"

func main() {
	var preco float64
	var resposta float64	

  fmt.Print("Digite o preço base do carro: ")
	fmt.Scan(&preco)

	fmt.Println("Digite 0 para NÃO e 1 para SIM")
	fmt.Print("O carro terá Ar-Condicionado? ")
	fmt.Scan(&resposta)
	if resposta == 1 { preco += 1750 }

	fmt.Print("O carro terá Pintura Metálica? ")
	fmt.Scan(&resposta)
	if resposta == 1 { preco += 800 }

	fmt.Print("O carro terá Vidro Elétrico? ")
	fmt.Scan(&resposta)
	if resposta == 1 { preco += 1200 }	

	fmt.Print("O carro terá Direção Hidraulica? ")
	fmt.Scan(&resposta)
	if resposta == 1 { preco += 2000 }

	fmt.Print("O carro ficou no preço: R$ ", preco)
}
