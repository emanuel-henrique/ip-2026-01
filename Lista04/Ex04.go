package main

import (
	"fmt"
)

func main() {
	var A [10]int
	
	contagem := make(map[int]int)

	fmt.Println("Digite 10 números inteiros:")

	for i := range 10 {
		fmt.Scan(&A[i])
		
		contagem[A[i]]++
	}

	fmt.Println("\n--- Elementos Repetidos ---")

	houveRepeticao := false
	for numero, qtd := range contagem {
		if qtd > 1 {
			fmt.Printf("O número %d se repete %d vezes.\n", numero, qtd)
			houveRepeticao = true
		}
	}

	if !houveRepeticao {
		fmt.Println("Não foram encontrados elementos repetidos.")
	}
}
