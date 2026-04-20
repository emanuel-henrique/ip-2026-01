package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Digite um número inteiro positivo N: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("O número deve ser positivo.")
		return
	}

	soma := 0
	for i := 1; i <= N; i++ {
		soma += i
	}

	fmt.Printf("O somatório de 1 a %d é: %d\n", N, soma)
}
