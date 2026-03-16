package main

import "fmt"

// Soma de progressão aritmética
func main() {
	var a1, r, n int

	fmt.Print("Digite os valores de a1, r, n (Nesta Ordem): ")
	fmt.Scan(&a1, &r, &n)
	fmt.Print(somaDosTermos(a1, r, n))
}

func somaDosTermos(a1, r, n int) int{
	result := (n * (2 * a1 + (n-1) * r)) / 2
	return result
}