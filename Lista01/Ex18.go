package main

import "fmt"

// Somatório simples
func main() {
	var result float64
	var n float64

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&n)

	for i := 1.00; i <= n; i++{
		result = result + (1/i)
	}
	
	fmt.Printf("%.6f", result)
}