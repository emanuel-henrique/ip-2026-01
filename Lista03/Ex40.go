package main

import "fmt"

func main() {
	valorIngresso := 6.00
	qtdIngresso := 130.00
	despesas := 300.00

	for valorIngresso > 1 {
		lucro := (qtdIngresso * valorIngresso) - despesas
		fmt.Printf("Valor do Ingresso: %.2f", valorIngresso)
		fmt.Print("   ")
		fmt.Printf("Ingressos Vendidos: %.2f", qtdIngresso)
		fmt.Print("   ")
		fmt.Printf("Lucro: %.2f\n", lucro)
		valorIngresso -= 0.60
		qtdIngresso += 30
	}
}
