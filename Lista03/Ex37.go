package main

import (
	"fmt"
	"strconv"
)

func main() {
	var octalStr string
	fmt.Print("Digite um número inteiro positivo na base 8: ")
	fmt.Scan(&octalStr)

	decimal, err := strconv.ParseInt(octalStr, 8, 64)
	if err != nil {
		fmt.Println("Erro: O valor digitado não é um número octal válido.")
		return
	}

	fmt.Printf("O número octal %s em decimal (base 10) é: %d\n", octalStr, decimal)
}
