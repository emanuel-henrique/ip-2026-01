package main

import (
	"fmt"
	"math"
)

func main() {
	var num, soma, somaPares, qtd, qtdPares, qtdImpares float64
	maior := -math.MaxFloat64
	menor := math.MaxFloat64

	for {
		fmt.Scan(&num)
		if num == 30000 {
			break
		}

		soma += num
		qtd++

		if num > maior { maior = num }
		if num < menor { menor = num }

		if int(num)%2 == 0 {
			somaPares += num
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	if qtd > 0 {
		fmt.Println("Soma:", soma)
		fmt.Println("Quantidade:", qtd)
		fmt.Printf("Média: %2.f\n", soma/qtd)
		fmt.Println("Maior:", maior)
		fmt.Println("Menor:", menor)
		if qtdPares > 0 {
			fmt.Printf("Média pares: %.2f\n", somaPares/qtdPares)
		}
		fmt.Printf("Porcentagem ímpares: %.2f%%\n", (qtdImpares/qtd)*100)
	}
}