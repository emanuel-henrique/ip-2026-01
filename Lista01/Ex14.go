package main

import "fmt"

// Quadrado de pares

func main() {
	var n int
	
	fmt.Print("Digite o número entre 5 a 2000: ")
	fmt.Scan(&n)

	if n < 5 || n > 2000 {
		fmt.Print("Digite um número válido")
	}

	for i := 1; i <= n; i++{
		if i%2 == 0 {
			fmt.Printf("%d ^ 2: %d\n", i, i*i)
		}
	}
}